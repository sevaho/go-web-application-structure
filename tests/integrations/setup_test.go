package integration

import (
	"testing"

	"github.com/sevaho/gowas/src/ingress/http"
	"github.com/sevaho/gowas/src/domain"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	httpIngress *http.HttpIngress
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	domain := domain.New()

	// attach HTTP ingress
	s.httpIngress = http.New(domain)

    // start HTTP server
	s.httpIngress.Serve()
}

func (s *TestSuite) SetupTest() {
}

func (s *TestSuite) TearDownTest() {
}

func (s *TestSuite) TearDownSuite() {
    // stop HTTP server
	s.httpIngress.ShutDown()
}
