package sqlpoolsconnectionpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsConnectionPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsConnectionPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsConnectionPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlpoolsconnectionpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsConnectionPoliciesClient: %+v", err)
	}

	return &SqlPoolsConnectionPoliciesClient{
		Client: client,
	}, nil
}
