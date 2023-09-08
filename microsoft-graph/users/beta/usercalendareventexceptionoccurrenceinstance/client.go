package usercalendareventexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventExceptionOccurrenceInstanceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &UserCalendarEventExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
