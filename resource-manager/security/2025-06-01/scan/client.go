package scan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScanClient struct {
	Client *resourcemanager.Client
}

func NewScanClientWithBaseURI(sdkApi sdkEnv.Api) (*ScanClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "scan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ScanClient: %+v", err)
	}

	return &ScanClient{
		Client: client,
	}, nil
}
