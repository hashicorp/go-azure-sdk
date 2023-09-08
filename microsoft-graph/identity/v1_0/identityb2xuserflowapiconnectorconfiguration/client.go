package identityb2xuserflowapiconnectorconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityB2xUserFlowApiConnectorConfigurationClient struct {
	Client *msgraph.Client
}

func NewIdentityB2xUserFlowApiConnectorConfigurationClientWithBaseURI(api sdkEnv.Api) (*IdentityB2xUserFlowApiConnectorConfigurationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityb2xuserflowapiconnectorconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityB2xUserFlowApiConnectorConfigurationClient: %+v", err)
	}

	return &IdentityB2xUserFlowApiConnectorConfigurationClient{
		Client: client,
	}, nil
}
