package majorversionupgradeprecheck

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MajorVersionUpgradePrecheckClient struct {
	Client *resourcemanager.Client
}

func NewMajorVersionUpgradePrecheckClientWithBaseURI(sdkApi sdkEnv.Api) (*MajorVersionUpgradePrecheckClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "majorversionupgradeprecheck", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MajorVersionUpgradePrecheckClient: %+v", err)
	}

	return &MajorVersionUpgradePrecheckClient{
		Client: client,
	}, nil
}
