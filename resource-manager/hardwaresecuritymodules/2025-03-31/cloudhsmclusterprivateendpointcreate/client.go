package cloudhsmclusterprivateendpointcreate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudHSMClusterPrivateEndpointCreateClient struct {
	Client *resourcemanager.Client
}

func NewCloudHSMClusterPrivateEndpointCreateClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudHSMClusterPrivateEndpointCreateClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cloudhsmclusterprivateendpointcreate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudHSMClusterPrivateEndpointCreateClient: %+v", err)
	}

	return &CloudHSMClusterPrivateEndpointCreateClient{
		Client: client,
	}, nil
}
