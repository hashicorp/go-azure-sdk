package b2xuserflowapiconnectorconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2xUserFlowApiConnectorConfigurationClient struct {
	Client *msgraph.Client
}

func NewB2xUserFlowApiConnectorConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*B2xUserFlowApiConnectorConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "b2xuserflowapiconnectorconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating B2xUserFlowApiConnectorConfigurationClient: %+v", err)
	}

	return &B2xUserFlowApiConnectorConfigurationClient{
		Client: client,
	}, nil
}
