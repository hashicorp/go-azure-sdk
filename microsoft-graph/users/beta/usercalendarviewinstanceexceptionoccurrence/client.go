package usercalendarviewinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarviewinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &UserCalendarViewInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
