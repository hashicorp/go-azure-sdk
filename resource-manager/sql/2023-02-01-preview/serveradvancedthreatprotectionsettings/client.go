package serveradvancedthreatprotectionsettings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerAdvancedThreatProtectionSettingsClient struct {
	Client *resourcemanager.Client
}

func NewServerAdvancedThreatProtectionSettingsClientWithBaseURI(sdkApi sdkEnv.Api) (*ServerAdvancedThreatProtectionSettingsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "serveradvancedthreatprotectionsettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerAdvancedThreatProtectionSettingsClient: %+v", err)
	}

	return &ServerAdvancedThreatProtectionSettingsClient{
		Client: client,
	}, nil
}
