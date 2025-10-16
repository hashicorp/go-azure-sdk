package v2025_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2025-06-01/antimalware"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2025-06-01/defenderforstorage"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2025-06-01/scan"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Antimalware        *antimalware.AntimalwareClient
	DefenderForStorage *defenderforstorage.DefenderForStorageClient
	Scan               *scan.ScanClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	antimalwareClient, err := antimalware.NewAntimalwareClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Antimalware client: %+v", err)
	}
	configureFunc(antimalwareClient.Client)

	defenderForStorageClient, err := defenderforstorage.NewDefenderForStorageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DefenderForStorage client: %+v", err)
	}
	configureFunc(defenderForStorageClient.Client)

	scanClient, err := scan.NewScanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Scan client: %+v", err)
	}
	configureFunc(scanClient.Client)

	return &Client{
		Antimalware:        antimalwareClient,
		DefenderForStorage: defenderForStorageClient,
		Scan:               scanClient,
	}, nil
}
