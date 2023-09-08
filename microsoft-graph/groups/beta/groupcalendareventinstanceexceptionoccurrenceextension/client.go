package groupcalendareventinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &GroupCalendarEventInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
