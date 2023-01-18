package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-infrastructure2/client-sms-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/stretchr/testify/assert"
)

type SmsClientFixtureV1 struct {
	Client version1.ISmsClientV1
}

func NewSmsClientFixtureV1(client version1.ISmsClientV1) *SmsClientFixtureV1 {
	return &SmsClientFixtureV1{
		Client: client,
	}
}

func (c *SmsClientFixtureV1) TestSendSmsToAddress(t *testing.T) {

	message := &version1.SmsMessageV1{
		To:   "+15203452335",
		Text: "{{text}}",
	}

	parameters := config.NewConfigParamsFromTuples(
		"subject", "Test Sms To Address",
		"text", "This is just a test",
	)

	err := c.Client.SendMessage(context.Background(), "123", message, parameters)
	assert.Nil(t, err)
}

func (c *SmsClientFixtureV1) TestSendSmsToRecipients(t *testing.T) {

	message := &version1.SmsMessageV1{
		Text: "This is just a test",
	}

	recipient := &version1.SmsRecipientV1{
		Id:    "1",
		Phone: "+152023452335",
		Name:  "Test Recipient",
	}

	err := c.Client.SendMessageToRecipient(context.Background(), "123", recipient, message, nil)
	assert.Nil(t, err)
}
