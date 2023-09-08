package usercalendarcalendarviewexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarviewexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewExceptionOccurrenceClient: %+v", err)
	}

	return &UserCalendarCalendarViewExceptionOccurrenceClient{
		Client: client,
	}, nil
}
