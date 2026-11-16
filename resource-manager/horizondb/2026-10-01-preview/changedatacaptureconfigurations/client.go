package changedatacaptureconfigurations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCaptureConfigurationsClient struct {
	Client *resourcemanager.Client
}

func NewChangeDataCaptureConfigurationsClientWithBaseURI(sdkApi sdkEnv.Api) (*ChangeDataCaptureConfigurationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "changedatacaptureconfigurations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChangeDataCaptureConfigurationsClient: %+v", err)
	}

	return &ChangeDataCaptureConfigurationsClient{
		Client: client,
	}, nil
}
