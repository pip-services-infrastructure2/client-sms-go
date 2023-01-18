package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type ISmsClientV1 interface {
	SendMessage(ctx context.Context, correlationId string, message *SmsMessageV1, parameters *config.ConfigParams) error

	SendMessageToRecipient(ctx context.Context, correlationId string, recipient *SmsRecipientV1,
		message *SmsMessageV1, parameters *config.ConfigParams) error

	SendMessageToRecipients(ctx context.Context, correlationId string, recipients []*SmsRecipientV1,
		message *SmsMessageV1, parameters *config.ConfigParams) error
}
