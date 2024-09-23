package b2xuserflowuserflowidentityprovider

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2xUserFlowUserFlowIdentityProviderClient struct {
	Client *msgraph.Client
}

func NewB2xUserFlowUserFlowIdentityProviderClientWithBaseURI(sdkApi sdkEnv.Api) (*B2xUserFlowUserFlowIdentityProviderClient, error) {
	client, err := msgraph.NewClient(sdkApi, "b2xuserflowuserflowidentityprovider", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating B2xUserFlowUserFlowIdentityProviderClient: %+v", err)
	}

	return &B2xUserFlowUserFlowIdentityProviderClient{
		Client: client,
	}, nil
}
