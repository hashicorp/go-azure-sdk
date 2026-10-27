package v2026_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2026-10-01/externalauths"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2026-10-01/hcpopenshiftclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2026-10-01/hcpopenshiftversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2026-10-01/hcpoperatoridentityrolesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2026-10-01/nodepools"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ExternalAuths               *externalauths.ExternalAuthsClient
	HcpOpenShiftClusters        *hcpopenshiftclusters.HcpOpenShiftClustersClient
	HcpOpenShiftVersions        *hcpopenshiftversions.HcpOpenShiftVersionsClient
	HcpOperatorIdentityRoleSets *hcpoperatoridentityrolesets.HcpOperatorIdentityRoleSetsClient
	NodePools                   *nodepools.NodePoolsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	externalAuthsClient, err := externalauths.NewExternalAuthsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExternalAuths client: %+v", err)
	}
	configureFunc(externalAuthsClient.Client)

	hcpOpenShiftClustersClient, err := hcpopenshiftclusters.NewHcpOpenShiftClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HcpOpenShiftClusters client: %+v", err)
	}
	configureFunc(hcpOpenShiftClustersClient.Client)

	hcpOpenShiftVersionsClient, err := hcpopenshiftversions.NewHcpOpenShiftVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HcpOpenShiftVersions client: %+v", err)
	}
	configureFunc(hcpOpenShiftVersionsClient.Client)

	hcpOperatorIdentityRoleSetsClient, err := hcpoperatoridentityrolesets.NewHcpOperatorIdentityRoleSetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HcpOperatorIdentityRoleSets client: %+v", err)
	}
	configureFunc(hcpOperatorIdentityRoleSetsClient.Client)

	nodePoolsClient, err := nodepools.NewNodePoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NodePools client: %+v", err)
	}
	configureFunc(nodePoolsClient.Client)

	return &Client{
		ExternalAuths:               externalAuthsClient,
		HcpOpenShiftClusters:        hcpOpenShiftClustersClient,
		HcpOpenShiftVersions:        hcpOpenShiftVersionsClient,
		HcpOperatorIdentityRoleSets: hcpOperatorIdentityRoleSetsClient,
		NodePools:                   nodePoolsClient,
	}, nil
}
