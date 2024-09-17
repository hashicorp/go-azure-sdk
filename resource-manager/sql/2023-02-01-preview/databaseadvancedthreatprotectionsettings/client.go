package databaseadvancedthreatprotectionsettings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseAdvancedThreatProtectionSettingsClient struct {
	Client *resourcemanager.Client
}

func NewDatabaseAdvancedThreatProtectionSettingsClientWithBaseURI(sdkApi sdkEnv.Api) (*DatabaseAdvancedThreatProtectionSettingsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "databaseadvancedthreatprotectionsettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatabaseAdvancedThreatProtectionSettingsClient: %+v", err)
	}

	return &DatabaseAdvancedThreatProtectionSettingsClient{
		Client: client,
	}, nil
}
