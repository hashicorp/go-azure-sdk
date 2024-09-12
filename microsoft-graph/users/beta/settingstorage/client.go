package settingstorage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingStorageClient struct {
	Client *msgraph.Client
}

func NewSettingStorageClientWithBaseURI(sdkApi sdkEnv.Api) (*SettingStorageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "settingstorage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SettingStorageClient: %+v", err)
	}

	return &SettingStorageClient{
		Client: client,
	}, nil
}
