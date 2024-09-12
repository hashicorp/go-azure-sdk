package permissionsanalyticgcppermissionscreepindexdistributionauthorizationsystem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalyticGcpPermissionsCreepIndexDistributionAuthorizationSystemClient struct {
	Client *msgraph.Client
}

func NewPermissionsAnalyticGcpPermissionsCreepIndexDistributionAuthorizationSystemClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionsAnalyticGcpPermissionsCreepIndexDistributionAuthorizationSystemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissionsanalyticgcppermissionscreepindexdistributionauthorizationsystem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionsAnalyticGcpPermissionsCreepIndexDistributionAuthorizationSystemClient: %+v", err)
	}

	return &PermissionsAnalyticGcpPermissionsCreepIndexDistributionAuthorizationSystemClient{
		Client: client,
	}, nil
}
