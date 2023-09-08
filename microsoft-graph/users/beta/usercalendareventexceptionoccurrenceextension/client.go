package usercalendareventexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &UserCalendarEventExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
