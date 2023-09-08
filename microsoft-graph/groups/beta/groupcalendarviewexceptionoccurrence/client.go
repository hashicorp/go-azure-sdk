package groupcalendarviewexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarViewExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarViewExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarViewExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarviewexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarViewExceptionOccurrenceClient: %+v", err)
	}

	return &GroupCalendarViewExceptionOccurrenceClient{
		Client: client,
	}, nil
}
