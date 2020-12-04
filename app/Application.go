package app

import (
	"errors"
	"html/template"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"main/server"
	"main/utils"

	database "g.ghn.vn/scte-common/godal"

	"main/apis/employeepb"
	
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
)

var (
	echoInstance *echo.Echo
	AppConfig    *utils.HoconConfig
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (r *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := r.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func getEchoInstance() *echo.Echo {
	if echoInstance == nil {
		echoInstance = echo.New()

		// Define middleware for API
		AppMiddleware(echoInstance)

		// Define route for API
		Routes(echoInstance)
	}

	return echoInstance
}

func initDatabase() {
	isUsePostgres := AppConfig.Conf.GetBoolean("postgresql.is_use", false)

	if isUsePostgres == true && utils.Global[utils.POSTGRES_ENTITY] == nil {
		host := AppConfig.Conf.GetString("postgresql.host")
		port := AppConfig.Conf.GetString("postgresql.port")
		dbname := AppConfig.Conf.GetString("postgresql.dbname")
		user := AppConfig.Conf.GetString("postgresql.user")
		pass := AppConfig.Conf.GetString("postgresql.pass")
		maxIdle := AppConfig.Conf.GetInt32("postgresql.max_idle_conn")
		maxOpen := AppConfig.Conf.GetInt32("postgresql.max_open_conn")

		p := database.Postgres{
			Host:        host,
			Port:        port,
			Dbname:      dbname,
			User:        user,
			Pass:        pass,
			MaxIdleConn: maxIdle,
			MaxOpenConn: maxOpen,
		}

		p.Connect()
		utils.Global[utils.POSTGRES_ENTITY] = p
	}
}

func startRestServer() {
	e := getEchoInstance()

	listenAddr := AppConfig.Conf.GetString("rest.listen_addr", "127.0.0.1")
	listenPort := AppConfig.Conf.GetInt32("rest.listen_port", 8080)
	request_timeout := AppConfig.Conf.GetInt32("api.request_timeout", 10)

	e.Static("/", "public")

	templates := make(map[string]*template.Template)
	templates["employee.html"] = template.Must(template.ParseFiles("public/views/employee.html", "public/views/base.html"))
	templates["department.html"] = template.Must(template.ParseFiles("public/views/department.html", "public/views/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	s := &http.Server{
		Addr:         listenAddr + ":" + strconv.Itoa(int(listenPort)),
		ReadTimeout:  time.Duration(request_timeout) * time.Second,
		WriteTimeout: time.Duration(request_timeout) * time.Second,
	}
	go e.Logger.Fatal(e.StartServer(s))
}

func startGrpcServer() {
	grpcPort := AppConfig.Conf.GetString("rest.address", "127.0.0.1:9090")
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// usersServer := users.UserServer{}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	// users.RegisterUsersServer(grpcServer, &usersServer)
	// pb.RegisterPApiServiceServer(grpcServer, newGrpcGateway())
	employeepb.RegisterEmployeeServiceServer(grpcServer, &server.EmployeeServer{})
	log.Infof("Start gRPC server at port: %s", grpcPort)
	go grpcServer.Serve(lis)
}

func initAPIServer() {
	restIsUse := AppConfig.Conf.GetBoolean("rest.is_use", false)
	grpcIsUse := AppConfig.Conf.GetBoolean("grpc.is_use", false)

	if grpcIsUse {
		startGrpcServer()
	}

	if restIsUse {
		startRestServer()
	}
}

func Start() {
	utils.InitAppConfig()
	AppConfig = utils.AppConfig
	log.Infof("Get app name [%s]", AppConfig.Conf.GetString("app.name"))

	// Init database
	initDatabase()

	// Init global for app
	AppInit(AppConfig)

	// Start server
	initAPIServer()
}
