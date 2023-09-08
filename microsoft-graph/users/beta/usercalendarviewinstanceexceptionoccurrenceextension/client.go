package usercalendarviewinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarviewinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &UserCalendarViewInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
