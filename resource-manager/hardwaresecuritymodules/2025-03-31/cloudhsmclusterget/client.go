package cloudhsmclusterget

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudHSMClusterGetClient struct {
	Client *resourcemanager.Client
}

func NewCloudHSMClusterGetClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudHSMClusterGetClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cloudhsmclusterget", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudHSMClusterGetClient: %+v", err)
	}

	return &CloudHSMClusterGetClient{
		Client: client,
	}, nil
}
