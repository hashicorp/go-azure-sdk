package hypervjobs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVJobsClient struct {
	Client *resourcemanager.Client
}

func NewHyperVJobsClientWithBaseURI(sdkApi sdkEnv.Api) (*HyperVJobsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "hypervjobs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HyperVJobsClient: %+v", err)
	}

	return &HyperVJobsClient{
		Client: client,
	}, nil
}
