package changedatacapturecapabilities

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCaptureCapabilitiesClient struct {
	Client *resourcemanager.Client
}

func NewChangeDataCaptureCapabilitiesClientWithBaseURI(sdkApi sdkEnv.Api) (*ChangeDataCaptureCapabilitiesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "changedatacapturecapabilities", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChangeDataCaptureCapabilitiesClient: %+v", err)
	}

	return &ChangeDataCaptureCapabilitiesClient{
		Client: client,
	}, nil
}
