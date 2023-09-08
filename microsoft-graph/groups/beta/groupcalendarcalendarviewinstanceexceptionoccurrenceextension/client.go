package groupcalendarcalendarviewinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &GroupCalendarCalendarViewInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
