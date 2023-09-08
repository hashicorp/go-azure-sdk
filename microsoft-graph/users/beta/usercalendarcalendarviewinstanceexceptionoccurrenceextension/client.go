package usercalendarcalendarviewinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarviewinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &UserCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
