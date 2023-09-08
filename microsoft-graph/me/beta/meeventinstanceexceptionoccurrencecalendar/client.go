package meeventinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeEventInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &MeEventInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
