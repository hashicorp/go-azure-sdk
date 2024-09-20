package hcireports

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HCIReportsClient struct {
	Client *resourcemanager.Client
}

func NewHCIReportsClientWithBaseURI(sdkApi sdkEnv.Api) (*HCIReportsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "hcireports", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HCIReportsClient: %+v", err)
	}

	return &HCIReportsClient{
		Client: client,
	}, nil
}
