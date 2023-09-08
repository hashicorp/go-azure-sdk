package usercalendarcalendarviewexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarviewexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &UserCalendarCalendarViewExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
