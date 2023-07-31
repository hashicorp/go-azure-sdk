package signupsettings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignUpSettingsClient struct {
	Client *resourcemanager.Client
}

func NewSignUpSettingsClientWithBaseURI(api environments.Api) (*SignUpSettingsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "signupsettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SignUpSettingsClient: %+v", err)
	}

	return &SignUpSettingsClient{
		Client: client,
	}, nil
}
