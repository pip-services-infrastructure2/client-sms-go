package build

import (
	clients1 "github.com/pip-services-infrastructure2/client-sms-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type SmsServiceFactory struct {
	*cbuild.Factory
}

func NewSmsServiceFactory() *SmsServiceFactory {
	c := &SmsServiceFactory{
		Factory: cbuild.NewFactory(),
	}

	cmdHttpClientDescriptor := cref.NewDescriptor("service-sms", "client", "commandable-http", "*", "1.0")
	nullClientDescriptor := cref.NewDescriptor("service-sms", "client", "null", "*", "1.0")

	c.RegisterType(cmdHttpClientDescriptor, clients1.NewSmsCommandableHttpClientV1)
	c.RegisterType(nullClientDescriptor, clients1.NewSmsNullClientV1)

	return c
}
