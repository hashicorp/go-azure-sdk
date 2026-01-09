package cloudpcresourcenamespaceresourceaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCResourceNamespaceResourceActionClient struct {
	Client *msgraph.Client
}

func NewCloudPCResourceNamespaceResourceActionClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudPCResourceNamespaceResourceActionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudpcresourcenamespaceresourceaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudPCResourceNamespaceResourceActionClient: %+v", err)
	}

	return &CloudPCResourceNamespaceResourceActionClient{
		Client: client,
	}, nil
}
