package usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarviewinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
