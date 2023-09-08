package usercalendarviewexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarviewexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &UserCalendarViewExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
