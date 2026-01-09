package cloudpcresourcenamespaceresourceactionauthenticationcontext

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCResourceNamespaceResourceActionAuthenticationContextClient struct {
	Client *msgraph.Client
}

func NewCloudPCResourceNamespaceResourceActionAuthenticationContextClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudPCResourceNamespaceResourceActionAuthenticationContextClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudpcresourcenamespaceresourceactionauthenticationcontext", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudPCResourceNamespaceResourceActionAuthenticationContextClient: %+v", err)
	}

	return &CloudPCResourceNamespaceResourceActionAuthenticationContextClient{
		Client: client,
	}, nil
}
