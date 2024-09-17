package managedrestorabledroppeddatabasebackupshorttermretentionpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "managedrestorabledroppeddatabasebackupshorttermretentionpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient: %+v", err)
	}

	return &ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient{
		Client: client,
	}, nil
}
