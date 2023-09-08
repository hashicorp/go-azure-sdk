package identityb2xuserflowuserflowidentityprovider

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityB2xUserFlowUserFlowIdentityProviderClient struct {
	Client *msgraph.Client
}

func NewIdentityB2xUserFlowUserFlowIdentityProviderClientWithBaseURI(api sdkEnv.Api) (*IdentityB2xUserFlowUserFlowIdentityProviderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityb2xuserflowuserflowidentityprovider", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityB2xUserFlowUserFlowIdentityProviderClient: %+v", err)
	}

	return &IdentityB2xUserFlowUserFlowIdentityProviderClient{
		Client: client,
	}, nil
}
