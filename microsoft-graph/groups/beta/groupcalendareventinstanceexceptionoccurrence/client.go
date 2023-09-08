package groupcalendareventinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventInstanceExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &GroupCalendarEventInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
