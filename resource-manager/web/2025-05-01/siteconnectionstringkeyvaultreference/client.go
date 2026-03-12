package siteconnectionstringkeyvaultreference

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteConnectionStringKeyVaultReferenceClient struct {
	Client *resourcemanager.Client
}

func NewSiteConnectionStringKeyVaultReferenceClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteConnectionStringKeyVaultReferenceClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "siteconnectionstringkeyvaultreference", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteConnectionStringKeyVaultReferenceClient: %+v", err)
	}

	return &SiteConnectionStringKeyVaultReferenceClient{
		Client: client,
	}, nil
}
