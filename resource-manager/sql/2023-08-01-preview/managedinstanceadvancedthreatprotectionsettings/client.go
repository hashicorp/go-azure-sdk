package managedinstanceadvancedthreatprotectionsettings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceAdvancedThreatProtectionSettingsClient struct {
	Client *resourcemanager.Client
}

func NewManagedInstanceAdvancedThreatProtectionSettingsClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedInstanceAdvancedThreatProtectionSettingsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "managedinstanceadvancedthreatprotectionsettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedInstanceAdvancedThreatProtectionSettingsClient: %+v", err)
	}

	return &ManagedInstanceAdvancedThreatProtectionSettingsClient{
		Client: client,
	}, nil
}
