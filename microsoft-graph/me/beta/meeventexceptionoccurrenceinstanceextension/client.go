package meeventexceptionoccurrenceinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventExceptionOccurrenceInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeEventExceptionOccurrenceInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventexceptionoccurrenceinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventExceptionOccurrenceInstanceExtensionClient: %+v", err)
	}

	return &MeEventExceptionOccurrenceInstanceExtensionClient{
		Client: client,
	}, nil
}
