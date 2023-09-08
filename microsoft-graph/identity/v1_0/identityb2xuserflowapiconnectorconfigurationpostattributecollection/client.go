package identityb2xuserflowapiconnectorconfigurationpostattributecollection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient struct {
	Client *msgraph.Client
}

func NewIdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClientWithBaseURI(api sdkEnv.Api) (*IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityb2xuserflowapiconnectorconfigurationpostattributecollection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient: %+v", err)
	}

	return &IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient{
		Client: client,
	}, nil
}
