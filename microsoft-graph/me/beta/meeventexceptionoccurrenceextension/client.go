package meeventexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewMeEventExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*MeEventExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &MeEventExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
