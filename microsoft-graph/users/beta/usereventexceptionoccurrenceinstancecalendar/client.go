package usereventexceptionoccurrenceinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventExceptionOccurrenceInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserEventExceptionOccurrenceInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventexceptionoccurrenceinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventExceptionOccurrenceInstanceCalendarClient: %+v", err)
	}

	return &UserEventExceptionOccurrenceInstanceCalendarClient{
		Client: client,
	}, nil
}
