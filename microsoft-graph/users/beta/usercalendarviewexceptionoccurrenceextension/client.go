package usercalendarviewexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarviewexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &UserCalendarViewExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
