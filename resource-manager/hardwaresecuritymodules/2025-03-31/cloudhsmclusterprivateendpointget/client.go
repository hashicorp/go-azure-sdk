package cloudhsmclusterprivateendpointget

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudHSMClusterPrivateEndpointGetClient struct {
	Client *resourcemanager.Client
}

func NewCloudHSMClusterPrivateEndpointGetClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudHSMClusterPrivateEndpointGetClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cloudhsmclusterprivateendpointget", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudHSMClusterPrivateEndpointGetClient: %+v", err)
	}

	return &CloudHSMClusterPrivateEndpointGetClient{
		Client: client,
	}, nil
}
