package usercalendarviewexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarviewexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewExceptionOccurrenceClient: %+v", err)
	}

	return &UserCalendarViewExceptionOccurrenceClient{
		Client: client,
	}, nil
}
