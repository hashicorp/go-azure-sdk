package eventexceptionoccurrenceinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventExceptionOccurrenceInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EventExceptionOccurrenceInstanceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventexceptionoccurrenceinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventExceptionOccurrenceInstanceExtensionClient: %+v", err)
	}

	return &EventExceptionOccurrenceInstanceExtensionClient{
		Client: client,
	}, nil
}
