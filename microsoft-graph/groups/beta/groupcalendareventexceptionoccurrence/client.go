package groupcalendareventexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventExceptionOccurrenceClient: %+v", err)
	}

	return &GroupCalendarEventExceptionOccurrenceClient{
		Client: client,
	}, nil
}
