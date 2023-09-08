package groupcalendarcalendarviewexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &GroupCalendarCalendarViewExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
