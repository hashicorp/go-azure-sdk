package permissionsanalyticawpermissionscreepindexdistribution

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalyticAwPermissionsCreepIndexDistributionClient struct {
	Client *msgraph.Client
}

func NewPermissionsAnalyticAwPermissionsCreepIndexDistributionClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionsAnalyticAwPermissionsCreepIndexDistributionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissionsanalyticawpermissionscreepindexdistribution", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionsAnalyticAwPermissionsCreepIndexDistributionClient: %+v", err)
	}

	return &PermissionsAnalyticAwPermissionsCreepIndexDistributionClient{
		Client: client,
	}, nil
}
