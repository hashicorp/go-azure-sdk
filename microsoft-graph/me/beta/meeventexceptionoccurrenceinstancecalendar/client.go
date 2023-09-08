package meeventexceptionoccurrenceinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventExceptionOccurrenceInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeEventExceptionOccurrenceInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventexceptionoccurrenceinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventExceptionOccurrenceInstanceCalendarClient: %+v", err)
	}

	return &MeEventExceptionOccurrenceInstanceCalendarClient{
		Client: client,
	}, nil
}
