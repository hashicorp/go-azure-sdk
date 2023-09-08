package identityb2xuserflowidentityprovider

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityB2xUserFlowIdentityProviderClient struct {
	Client *msgraph.Client
}

func NewIdentityB2xUserFlowIdentityProviderClientWithBaseURI(api sdkEnv.Api) (*IdentityB2xUserFlowIdentityProviderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityb2xuserflowidentityprovider", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityB2xUserFlowIdentityProviderClient: %+v", err)
	}

	return &IdentityB2xUserFlowIdentityProviderClient{
		Client: client,
	}, nil
}
