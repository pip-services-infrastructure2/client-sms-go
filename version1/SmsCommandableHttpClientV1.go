package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type SmsCommandableHttpClientV1 struct {
	defaultParameters *cconf.ConfigParams
	*clients.CommandableHttpClient
}

func NewSmsCommandableHttpClientV1() *SmsCommandableHttpClientV1 {

	return &SmsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/sms"),
	}
}

func NewSmsCommandableHttpClientV1WithParameters(config any) *SmsCommandableHttpClientV1 {
	thisConfig := cconf.NewConfigParamsFromValue(config)

	c := &SmsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/sms"),
	}

	c.defaultParameters = thisConfig.GetSection("parameters")

	if config != nil {
		c.Configure(context.Background(), thisConfig)
	}
	return c
}

func (c *SmsCommandableHttpClientV1) SendMessage(ctx context.Context, correlationId string, message *SmsMessageV1, parameters *cconf.ConfigParams) error {
	parameters = c.defaultParameters.Override(parameters)

	_, err := c.CallCommand(ctx, "send_message", correlationId, data.NewAnyValueMapFromTuples(
		"message", message,
		"parameters", parameters,
	))

	return err
}

func (c *SmsCommandableHttpClientV1) SendMessageToRecipient(ctx context.Context, correlationId string, recipient *SmsRecipientV1, message *SmsMessageV1, parameters *cconf.ConfigParams) error {
	parameters = c.defaultParameters.Override(parameters)

	_, err := c.CallCommand(ctx, "send_message_to_recipient", correlationId, data.NewAnyValueMapFromTuples(
		"recipient", recipient,
		"message", message,
		"parameters", parameters,
	))

	return err
}

func (c *SmsCommandableHttpClientV1) SendMessageToRecipients(ctx context.Context, correlationId string, recipients []*SmsRecipientV1, message *SmsMessageV1, parameters *cconf.ConfigParams) error {
	parameters = c.defaultParameters.Override(parameters)

	_, err := c.CallCommand(ctx, "send_message_to_recipients", correlationId, data.NewAnyValueMapFromTuples(
		"recipients", recipients,
		"message", message,
		"parameters", parameters,
	))

	return err
}
