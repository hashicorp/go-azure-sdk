package groupcalendarcalendarviewexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewExceptionOccurrenceClient: %+v", err)
	}

	return &GroupCalendarCalendarViewExceptionOccurrenceClient{
		Client: client,
	}, nil
}
