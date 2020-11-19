package test

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
}
