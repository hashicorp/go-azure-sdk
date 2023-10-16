package componentcontinuousexportapis

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComponentContinuousExportAPIsClient struct {
	Client *resourcemanager.Client
}

func NewComponentContinuousExportAPIsClientWithBaseURI(sdkApi sdkEnv.Api) (*ComponentContinuousExportAPIsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "componentcontinuousexportapis", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComponentContinuousExportAPIsClient: %+v", err)
	}

	return &ComponentContinuousExportAPIsClient{
		Client: client,
	}, nil
}
