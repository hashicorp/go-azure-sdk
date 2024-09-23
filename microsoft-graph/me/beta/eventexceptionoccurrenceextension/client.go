package eventexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EventExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &EventExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
