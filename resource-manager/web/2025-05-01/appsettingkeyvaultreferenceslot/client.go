package appsettingkeyvaultreferenceslot

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppSettingKeyVaultReferenceSlotClient struct {
	Client *resourcemanager.Client
}

func NewAppSettingKeyVaultReferenceSlotClientWithBaseURI(sdkApi sdkEnv.Api) (*AppSettingKeyVaultReferenceSlotClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "appsettingkeyvaultreferenceslot", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppSettingKeyVaultReferenceSlotClient: %+v", err)
	}

	return &AppSettingKeyVaultReferenceSlotClient{
		Client: client,
	}, nil
}
