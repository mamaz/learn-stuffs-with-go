package simpletest

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type CacheSuite struct {
	suite.Suite
	redisC   testcontainers.Container
	endpoint string
}

func (s *CacheSuite) SetupTest() {
	log.Println("setup suite ....")
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	var err error
	s.redisC, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		s.T().Error(err)
	}

	var redisConnErr error
	s.endpoint, redisConnErr = s.redisC.Endpoint(context.Background(), "")
	if redisConnErr != nil {
		s.T().Error(redisConnErr)
	}
}

func (s *CacheSuite) TestSetCache() {
	cache := NewCache(s.endpoint)
	cache.SetCache("test", "tist")
	val, err := cache.GetCache("test")

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "tist", val)
}

func (s *CacheSuite) TestGetCache() {
	cache := NewCache(s.endpoint)
	cache.SetCache("test", "tist")
	val, err := cache.GetCache("test")

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), "tist", val)
}

func (s *CacheSuite) TearDownSuite() {
	log.Println("tear down suite ....")
	if err := s.redisC.Terminate(context.Background()); err != nil {
		s.T().Error(err)
	}
}

func TestMyTestSuite(t *testing.T) {
	suite.Run(t, new(CacheSuite))
}
