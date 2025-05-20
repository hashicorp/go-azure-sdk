package cloudhsmclusterprivateendpointdelete

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudHSMClusterPrivateEndpointDeleteClient struct {
	Client *resourcemanager.Client
}

func NewCloudHSMClusterPrivateEndpointDeleteClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudHSMClusterPrivateEndpointDeleteClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cloudhsmclusterprivateendpointdelete", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudHSMClusterPrivateEndpointDeleteClient: %+v", err)
	}

	return &CloudHSMClusterPrivateEndpointDeleteClient{
		Client: client,
	}, nil
}
