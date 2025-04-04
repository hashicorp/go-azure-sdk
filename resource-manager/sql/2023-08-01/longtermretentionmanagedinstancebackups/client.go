package longtermretentionmanagedinstancebackups

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LongTermRetentionManagedInstanceBackupsClient struct {
	Client *resourcemanager.Client
}

func NewLongTermRetentionManagedInstanceBackupsClientWithBaseURI(sdkApi sdkEnv.Api) (*LongTermRetentionManagedInstanceBackupsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "longtermretentionmanagedinstancebackups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LongTermRetentionManagedInstanceBackupsClient: %+v", err)
	}

	return &LongTermRetentionManagedInstanceBackupsClient{
		Client: client,
	}, nil
}
