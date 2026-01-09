package settingstoragequota

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingStorageQuotaClient struct {
	Client *msgraph.Client
}

func NewSettingStorageQuotaClientWithBaseURI(sdkApi sdkEnv.Api) (*SettingStorageQuotaClient, error) {
	client, err := msgraph.NewClient(sdkApi, "settingstoragequota", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SettingStorageQuotaClient: %+v", err)
	}

	return &SettingStorageQuotaClient{
		Client: client,
	}, nil
}
