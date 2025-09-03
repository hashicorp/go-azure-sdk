package groupquotalocationsettings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupQuotaLocationSettingsClient struct {
	Client *resourcemanager.Client
}

func NewGroupQuotaLocationSettingsClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupQuotaLocationSettingsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "groupquotalocationsettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupQuotaLocationSettingsClient: %+v", err)
	}

	return &GroupQuotaLocationSettingsClient{
		Client: client,
	}, nil
}
