package b2xuserflowidentityprovider

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2xUserFlowIdentityProviderClient struct {
	Client *msgraph.Client
}

func NewB2xUserFlowIdentityProviderClientWithBaseURI(sdkApi sdkEnv.Api) (*B2xUserFlowIdentityProviderClient, error) {
	client, err := msgraph.NewClient(sdkApi, "b2xuserflowidentityprovider", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating B2xUserFlowIdentityProviderClient: %+v", err)
	}

	return &B2xUserFlowIdentityProviderClient{
		Client: client,
	}, nil
}
