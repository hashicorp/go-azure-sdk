package longtermretentionbackup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LongTermRetentionBackupClient struct {
	Client *resourcemanager.Client
}

func NewLongTermRetentionBackupClientWithBaseURI(sdkApi sdkEnv.Api) (*LongTermRetentionBackupClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "longtermretentionbackup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LongTermRetentionBackupClient: %+v", err)
	}

	return &LongTermRetentionBackupClient{
		Client: client,
	}, nil
}
