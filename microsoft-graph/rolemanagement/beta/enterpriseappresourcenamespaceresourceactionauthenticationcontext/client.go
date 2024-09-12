package enterpriseappresourcenamespaceresourceactionauthenticationcontext

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppResourceNamespaceResourceActionAuthenticationContextClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppResourceNamespaceResourceActionAuthenticationContextClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseappresourcenamespaceresourceactionauthenticationcontext", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppResourceNamespaceResourceActionAuthenticationContextClient: %+v", err)
	}

	return &EnterpriseAppResourceNamespaceResourceActionAuthenticationContextClient{
		Client: client,
	}, nil
}
