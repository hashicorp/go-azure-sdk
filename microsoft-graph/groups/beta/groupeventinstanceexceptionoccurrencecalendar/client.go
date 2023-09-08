package groupeventinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupEventInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &GroupEventInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
