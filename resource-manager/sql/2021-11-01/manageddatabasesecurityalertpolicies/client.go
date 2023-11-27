package manageddatabasesecurityalertpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseSecurityAlertPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewManagedDatabaseSecurityAlertPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDatabaseSecurityAlertPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "manageddatabasesecurityalertpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDatabaseSecurityAlertPoliciesClient: %+v", err)
	}

	return &ManagedDatabaseSecurityAlertPoliciesClient{
		Client: client,
	}, nil
}
