package sqlpoolsgeobackuppolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsGeoBackupPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsGeoBackupPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsGeoBackupPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlpoolsgeobackuppolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsGeoBackupPoliciesClient: %+v", err)
	}

	return &SqlPoolsGeoBackupPoliciesClient{
		Client: client,
	}, nil
}
