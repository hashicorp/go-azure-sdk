package cloudpcresourcenamespaceresourceactionresourcescope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCResourceNamespaceResourceActionResourceScopeClient struct {
	Client *msgraph.Client
}

func NewCloudPCResourceNamespaceResourceActionResourceScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudPCResourceNamespaceResourceActionResourceScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudpcresourcenamespaceresourceactionresourcescope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudPCResourceNamespaceResourceActionResourceScopeClient: %+v", err)
	}

	return &CloudPCResourceNamespaceResourceActionResourceScopeClient{
		Client: client,
	}, nil
}
