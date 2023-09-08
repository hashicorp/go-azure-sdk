package groupcalendareventexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventExceptionOccurrenceInstanceClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &GroupCalendarEventExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
