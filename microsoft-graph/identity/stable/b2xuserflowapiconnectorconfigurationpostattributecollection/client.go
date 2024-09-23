package b2xuserflowapiconnectorconfigurationpostattributecollection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient struct {
	Client *msgraph.Client
}

func NewB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClientWithBaseURI(sdkApi sdkEnv.Api) (*B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "b2xuserflowapiconnectorconfigurationpostattributecollection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient: %+v", err)
	}

	return &B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient{
		Client: client,
	}, nil
}
