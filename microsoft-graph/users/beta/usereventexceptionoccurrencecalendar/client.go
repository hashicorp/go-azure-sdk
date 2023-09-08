package usereventexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserEventExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserEventExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &UserEventExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
