package v2025_07_25

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2025-07-25/openshiftclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2025-07-25/openshiftversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2025-07-25/platformworkloadidentityrolesets"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	OpenShiftClusters                *openshiftclusters.OpenShiftClustersClient
	OpenShiftVersions                *openshiftversions.OpenShiftVersionsClient
	PlatformWorkloadIdentityRoleSets *platformworkloadidentityrolesets.PlatformWorkloadIdentityRoleSetsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	openShiftClustersClient, err := openshiftclusters.NewOpenShiftClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OpenShiftClusters client: %+v", err)
	}
	configureFunc(openShiftClustersClient.Client)

	openShiftVersionsClient, err := openshiftversions.NewOpenShiftVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OpenShiftVersions client: %+v", err)
	}
	configureFunc(openShiftVersionsClient.Client)

	platformWorkloadIdentityRoleSetsClient, err := platformworkloadidentityrolesets.NewPlatformWorkloadIdentityRoleSetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlatformWorkloadIdentityRoleSets client: %+v", err)
	}
	configureFunc(platformWorkloadIdentityRoleSetsClient.Client)

	return &Client{
		OpenShiftClusters:                openShiftClustersClient,
		OpenShiftVersions:                openShiftVersionsClient,
		PlatformWorkloadIdentityRoleSets: platformWorkloadIdentityRoleSetsClient,
	}, nil
}
