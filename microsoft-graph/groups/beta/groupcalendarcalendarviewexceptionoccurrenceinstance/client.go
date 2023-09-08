package groupcalendarcalendarviewexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &GroupCalendarCalendarViewExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
