package groupcalendareventexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &GroupCalendarEventExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
