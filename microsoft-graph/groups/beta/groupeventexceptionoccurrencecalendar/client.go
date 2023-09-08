package groupeventexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupEventExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupEventExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &GroupEventExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
