package usercalendareventinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventInstanceExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &UserCalendarEventInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
