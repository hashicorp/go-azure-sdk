package advancedthreatprotectionsettingsmodels

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionSettingsModelsClient struct {
	Client *resourcemanager.Client
}

func NewAdvancedThreatProtectionSettingsModelsClientWithBaseURI(sdkApi sdkEnv.Api) (*AdvancedThreatProtectionSettingsModelsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "advancedthreatprotectionsettingsmodels", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdvancedThreatProtectionSettingsModelsClient: %+v", err)
	}

	return &AdvancedThreatProtectionSettingsModelsClient{
		Client: client,
	}, nil
}
