package openshiftclusters

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftClustersClient struct {
	Client *resourcemanager.Client
}

func NewOpenShiftClustersClientWithBaseURI(api sdkEnv.Api) (*OpenShiftClustersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "openshiftclusters", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OpenShiftClustersClient: %+v", err)
	}

	return &OpenShiftClustersClient{
		Client: client,
	}, nil
}
