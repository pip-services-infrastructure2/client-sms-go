package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-infrastructure2/client-sms-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type smsCommandableHttpClientV1Test struct {
	client  *version1.SmsCommandableHttpClientV1
	fixture *SmsClientFixtureV1
}

func newSmsCommandableHttpClientV1Test() *smsCommandableHttpClientV1Test {
	return &smsCommandableHttpClientV1Test{}
}

func (c *smsCommandableHttpClientV1Test) setup(t *testing.T) *SmsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewSmsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewSmsClientFixtureV1(c.client)

	return c.fixture
}

func (c *smsCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestSendSmstoAddress(t *testing.T) {
	c := newSmsCommandableHttpClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestSendSmsToAddress(t)
}

func TestSendSmstoRecipients(t *testing.T) {
	c := newSmsCommandableHttpClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestSendSmsToRecipients(t)
}
