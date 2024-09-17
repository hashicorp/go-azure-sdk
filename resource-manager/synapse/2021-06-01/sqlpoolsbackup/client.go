package sqlpoolsbackup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsBackupClient struct {
	Client *resourcemanager.Client
}

func NewSqlPoolsBackupClientWithBaseURI(sdkApi sdkEnv.Api) (*SqlPoolsBackupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sqlpoolsbackup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsBackupClient: %+v", err)
	}

	return &SqlPoolsBackupClient{
		Client: client,
	}, nil
}
