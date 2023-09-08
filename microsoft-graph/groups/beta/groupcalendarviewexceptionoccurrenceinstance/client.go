package groupcalendarviewexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarViewExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarViewExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarviewexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarViewExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &GroupCalendarViewExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
