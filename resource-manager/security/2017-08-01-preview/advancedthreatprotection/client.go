package advancedthreatprotection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionClient struct {
	Client *resourcemanager.Client
}

func NewAdvancedThreatProtectionClientWithBaseURI(sdkApi sdkEnv.Api) (*AdvancedThreatProtectionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "advancedthreatprotection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdvancedThreatProtectionClient: %+v", err)
	}

	return &AdvancedThreatProtectionClient{
		Client: client,
	}, nil
}
