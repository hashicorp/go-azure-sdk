package eventinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*EventInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &EventInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
