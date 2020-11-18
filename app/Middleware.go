package app

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alexcesaro/statsd.v2"
)

type ProfilerConfig struct {
	Address     string
	ServiceName string
	Skipper     middleware.Skipper
}

func AppMiddleware(e *echo.Echo) {
	// Client cannot send request that exceeds this size
	maxBody := AppConfig.Conf.GetString("api.body_limit", "64KiB")
	e.Use(middleware.BodyLimit(maxBody))

	// Add Logger middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} \t ${remote_ip} \t ${host} \t ${method} \t ${uri} \t ${status} \t ${latency_human} \t ${latency} \n",
	}))

	// Push time to statsd
	statsdAddr := AppConfig.Conf.GetString("statsd.address")
	statsdServiceName := AppConfig.Conf.GetString("statsd.service_name")
	statsdConfig := ProfilerConfig{
		Address:     statsdAddr,
		ServiceName: statsdServiceName,
		Skipper:     defaultSkipper,
	}
	e.Use(ProfilerWithStatsD(statsdConfig))
}

func defaultSkipper(c echo.Context) bool {
	return false
}

func ProfilerWithStatsD(conf ProfilerConfig) echo.MiddlewareFunc {
	client, err := statsd.New(statsd.Address(conf.Address))
	if err != nil {
		panic(fmt.Sprintf("Cannot connect with statsd %s", conf.Address))
	}
	log.Infof("Connected to statsd %s successful", conf.Address)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if conf.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			t := client.NewTiming()
			next(c)
			s := strings.ToLower(fmt.Sprintf("response.%s.%s.%s.%d", conf.ServiceName, req.Method, req.URL.Path, res.Status))
			t.Send(s)

			return
		}
	}
}
