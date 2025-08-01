package accountconnectionresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountConnectionResourceClient struct {
	Client *resourcemanager.Client
}

func NewAccountConnectionResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*AccountConnectionResourceClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "accountconnectionresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccountConnectionResourceClient: %+v", err)
	}

	return &AccountConnectionResourceClient{
		Client: client,
	}, nil
}
