package settingstoragequotaservice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingStorageQuotaServiceClient struct {
	Client *msgraph.Client
}

func NewSettingStorageQuotaServiceClientWithBaseURI(sdkApi sdkEnv.Api) (*SettingStorageQuotaServiceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "settingstoragequotaservice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SettingStorageQuotaServiceClient: %+v", err)
	}

	return &SettingStorageQuotaServiceClient{
		Client: client,
	}, nil
}
