package signaldefinitions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignalDefinitionsClient struct {
	Client *resourcemanager.Client
}

func NewSignalDefinitionsClientWithBaseURI(sdkApi sdkEnv.Api) (*SignalDefinitionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "signaldefinitions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SignalDefinitionsClient: %+v", err)
	}

	return &SignalDefinitionsClient{
		Client: client,
	}, nil
}
