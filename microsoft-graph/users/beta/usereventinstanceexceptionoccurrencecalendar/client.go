package usereventinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserEventInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &UserEventInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
