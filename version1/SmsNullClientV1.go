package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type SmsNullClientV1 struct {
}

func NewSmsNullClientV1() *SmsNullClientV1 {
	return &SmsNullClientV1{}
}

func (c *SmsNullClientV1) SendMessage(ctx context.Context, correlationId string, message *SmsMessageV1, parameters *config.ConfigParams) error {
	return nil
}

func (c *SmsNullClientV1) SendMessageToRecipient(ctx context.Context, correlationId string, recipient *SmsRecipientV1, message *SmsMessageV1, parameters *config.ConfigParams) error {
	return nil
}

func (c *SmsNullClientV1) SendMessageToRecipients(ctx context.Context, correlationId string, recipients []*SmsRecipientV1, message *SmsMessageV1, parameters *config.ConfigParams) error {
	return nil
}
