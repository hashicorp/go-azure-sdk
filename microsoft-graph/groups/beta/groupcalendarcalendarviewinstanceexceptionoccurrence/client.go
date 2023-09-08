package groupcalendarcalendarviewinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &GroupCalendarCalendarViewInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
