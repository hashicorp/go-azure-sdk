package usercalendarcalendarviewinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarviewinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &UserCalendarCalendarViewInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
