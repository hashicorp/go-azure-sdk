package groupcalendarviewexceptionoccurrenceinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarViewExceptionOccurrenceInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarViewExceptionOccurrenceInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarviewexceptionoccurrenceinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarViewExceptionOccurrenceInstanceExtensionClient: %+v", err)
	}

	return &GroupCalendarViewExceptionOccurrenceInstanceExtensionClient{
		Client: client,
	}, nil
}
