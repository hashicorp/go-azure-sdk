package defenderresourcenamespaceresourceactionauthenticationcontext

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderResourceNamespaceResourceActionAuthenticationContextClient struct {
	Client *msgraph.Client
}

func NewDefenderResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi sdkEnv.Api) (*DefenderResourceNamespaceResourceActionAuthenticationContextClient, error) {
	client, err := msgraph.NewClient(sdkApi, "defenderresourcenamespaceresourceactionauthenticationcontext", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DefenderResourceNamespaceResourceActionAuthenticationContextClient: %+v", err)
	}

	return &DefenderResourceNamespaceResourceActionAuthenticationContextClient{
		Client: client,
	}, nil
}
