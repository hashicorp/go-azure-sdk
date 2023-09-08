package usercalendarcalendarviewexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarviewexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &UserCalendarCalendarViewExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
