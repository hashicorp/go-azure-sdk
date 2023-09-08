package identityb2xuserflowapiconnectorconfigurationpostfederationsignup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient struct {
	Client *msgraph.Client
}

func NewIdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClientWithBaseURI(api sdkEnv.Api) (*IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityb2xuserflowapiconnectorconfigurationpostfederationsignup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient: %+v", err)
	}

	return &IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient{
		Client: client,
	}, nil
}
