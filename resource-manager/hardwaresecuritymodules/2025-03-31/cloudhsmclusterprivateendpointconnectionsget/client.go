package cloudhsmclusterprivateendpointconnectionsget

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudHSMClusterPrivateEndpointConnectionsGetClient struct {
	Client *resourcemanager.Client
}

func NewCloudHSMClusterPrivateEndpointConnectionsGetClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudHSMClusterPrivateEndpointConnectionsGetClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cloudhsmclusterprivateendpointconnectionsget", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudHSMClusterPrivateEndpointConnectionsGetClient: %+v", err)
	}

	return &CloudHSMClusterPrivateEndpointConnectionsGetClient{
		Client: client,
	}, nil
}
