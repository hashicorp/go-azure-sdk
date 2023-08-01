package backuplongtermretentionpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupLongTermRetentionPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewBackupLongTermRetentionPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*BackupLongTermRetentionPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "backuplongtermretentionpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BackupLongTermRetentionPoliciesClient: %+v", err)
	}

	return &BackupLongTermRetentionPoliciesClient{
		Client: client,
	}, nil
}
