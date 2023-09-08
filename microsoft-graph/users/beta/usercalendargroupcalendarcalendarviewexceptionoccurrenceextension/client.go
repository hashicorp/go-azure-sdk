package usercalendargroupcalendarcalendarviewexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarviewexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
