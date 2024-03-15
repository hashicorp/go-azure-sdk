package securityconnectors

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConnectorsClient struct {
	Client *resourcemanager.Client
}

func NewSecurityConnectorsClientWithBaseURI(sdkApi sdkEnv.Api) (*SecurityConnectorsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "securityconnectors", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SecurityConnectorsClient: %+v", err)
	}

	return &SecurityConnectorsClient{
		Client: client,
	}, nil
}
