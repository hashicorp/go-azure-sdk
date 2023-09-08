package usercalendargroupcalendarcalendarviewexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarviewexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
