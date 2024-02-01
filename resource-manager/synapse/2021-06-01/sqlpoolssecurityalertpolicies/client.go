package sqlpoolssecurityalertpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsSecurityAlertPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsSecurityAlertPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsSecurityAlertPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sqlpoolssecurityalertpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsSecurityAlertPoliciesClient: %+v", err)
	}

	return &SqlPoolsSecurityAlertPoliciesClient{
		Client: client,
	}, nil
}
