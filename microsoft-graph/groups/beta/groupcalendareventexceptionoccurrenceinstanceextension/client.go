package groupcalendareventexceptionoccurrenceinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventExceptionOccurrenceInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventExceptionOccurrenceInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventexceptionoccurrenceinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventExceptionOccurrenceInstanceExtensionClient: %+v", err)
	}

	return &GroupCalendarEventExceptionOccurrenceInstanceExtensionClient{
		Client: client,
	}, nil
}
