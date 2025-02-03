package volumesonpremmigrationfinalize

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VolumesOnPremMigrationFinalizeClient struct {
	Client *resourcemanager.Client
}

func NewVolumesOnPremMigrationFinalizeClientWithBaseURI(sdkApi sdkEnv.Api) (*VolumesOnPremMigrationFinalizeClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "volumesonpremmigrationfinalize", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VolumesOnPremMigrationFinalizeClient: %+v", err)
	}

	return &VolumesOnPremMigrationFinalizeClient{
		Client: client,
	}, nil
}
