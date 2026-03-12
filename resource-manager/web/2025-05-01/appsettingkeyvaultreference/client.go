package appsettingkeyvaultreference

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppSettingKeyVaultReferenceClient struct {
	Client *resourcemanager.Client
}

func NewAppSettingKeyVaultReferenceClientWithBaseURI(sdkApi sdkEnv.Api) (*AppSettingKeyVaultReferenceClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "appsettingkeyvaultreference", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppSettingKeyVaultReferenceClient: %+v", err)
	}

	return &AppSettingKeyVaultReferenceClient{
		Client: client,
	}, nil
}
