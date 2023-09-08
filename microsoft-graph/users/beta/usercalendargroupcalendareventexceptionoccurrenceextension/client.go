package usercalendargroupcalendareventexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendareventexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
