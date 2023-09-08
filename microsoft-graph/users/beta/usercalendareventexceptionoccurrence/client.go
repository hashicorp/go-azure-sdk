package usercalendareventexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventExceptionOccurrenceClient: %+v", err)
	}

	return &UserCalendarEventExceptionOccurrenceClient{
		Client: client,
	}, nil
}
