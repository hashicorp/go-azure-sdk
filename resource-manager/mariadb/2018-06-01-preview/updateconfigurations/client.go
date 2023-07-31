package updateconfigurations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateConfigurationsClient struct {
	Client *resourcemanager.Client
}

func NewUpdateConfigurationsClientWithBaseURI(api sdkEnv.Api) (*UpdateConfigurationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "updateconfigurations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UpdateConfigurationsClient: %+v", err)
	}

	return &UpdateConfigurationsClient{
		Client: client,
	}, nil
}
