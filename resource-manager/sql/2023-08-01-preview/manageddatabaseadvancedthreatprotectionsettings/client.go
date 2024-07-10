package manageddatabaseadvancedthreatprotectionsettings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseAdvancedThreatProtectionSettingsClient struct {
	Client *resourcemanager.Client
}

func NewManagedDatabaseAdvancedThreatProtectionSettingsClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedDatabaseAdvancedThreatProtectionSettingsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "manageddatabaseadvancedthreatprotectionsettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDatabaseAdvancedThreatProtectionSettingsClient: %+v", err)
	}

	return &ManagedDatabaseAdvancedThreatProtectionSettingsClient{
		Client: client,
	}, nil
}
