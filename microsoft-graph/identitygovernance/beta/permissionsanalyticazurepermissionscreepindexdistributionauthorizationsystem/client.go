package permissionsanalyticazurepermissionscreepindexdistributionauthorizationsystem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalyticAzurePermissionsCreepIndexDistributionAuthorizationSystemClient struct {
	Client *msgraph.Client
}

func NewPermissionsAnalyticAzurePermissionsCreepIndexDistributionAuthorizationSystemClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionsAnalyticAzurePermissionsCreepIndexDistributionAuthorizationSystemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissionsanalyticazurepermissionscreepindexdistributionauthorizationsystem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionsAnalyticAzurePermissionsCreepIndexDistributionAuthorizationSystemClient: %+v", err)
	}

	return &PermissionsAnalyticAzurePermissionsCreepIndexDistributionAuthorizationSystemClient{
		Client: client,
	}, nil
}
