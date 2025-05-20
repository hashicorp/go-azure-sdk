package cloudhsmclusterprivatelinkget

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudHSMClusterPrivateLinkGetClient struct {
	Client *resourcemanager.Client
}

func NewCloudHSMClusterPrivateLinkGetClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudHSMClusterPrivateLinkGetClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cloudhsmclusterprivatelinkget", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudHSMClusterPrivateLinkGetClient: %+v", err)
	}

	return &CloudHSMClusterPrivateLinkGetClient{
		Client: client,
	}, nil
}
