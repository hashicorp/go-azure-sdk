package usercalendargroupcalendareventexceptionoccurrenceinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendareventexceptionoccurrenceinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventExceptionOccurrenceInstanceExtensionClient{
		Client: client,
	}, nil
}
