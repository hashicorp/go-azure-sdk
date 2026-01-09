package federationconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FederationConfigurationClient struct {
	Client *msgraph.Client
}

func NewFederationConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*FederationConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "federationconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FederationConfigurationClient: %+v", err)
	}

	return &FederationConfigurationClient{
		Client: client,
	}, nil
}
