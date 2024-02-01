package sqlpoolsdatamaskingpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsDataMaskingPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsDataMaskingPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsDataMaskingPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlpoolsdatamaskingpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsDataMaskingPoliciesClient: %+v", err)
	}

	return &SqlPoolsDataMaskingPoliciesClient{
		Client: client,
	}, nil
}
