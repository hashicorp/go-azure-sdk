package credits

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreditsClient struct {
	Client *resourcemanager.Client
}

func NewCreditsClientWithBaseURI(sdkApi sdkEnv.Api) (*CreditsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "credits", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CreditsClient: %+v", err)
	}

	return &CreditsClient{
		Client: client,
	}, nil
}
