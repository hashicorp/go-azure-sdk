package usercalendareventinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &UserCalendarEventInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
