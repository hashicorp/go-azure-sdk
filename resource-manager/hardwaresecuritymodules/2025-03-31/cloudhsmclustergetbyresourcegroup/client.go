package cloudhsmclustergetbyresourcegroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudHSMClusterGetByResourceGroupClient struct {
	Client *resourcemanager.Client
}

func NewCloudHSMClusterGetByResourceGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudHSMClusterGetByResourceGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cloudhsmclustergetbyresourcegroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudHSMClusterGetByResourceGroupClient: %+v", err)
	}

	return &CloudHSMClusterGetByResourceGroupClient{
		Client: client,
	}, nil
}
