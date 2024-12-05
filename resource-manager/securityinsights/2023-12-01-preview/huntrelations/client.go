package huntrelations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HuntRelationsClient struct {
	Client *resourcemanager.Client
}

func NewHuntRelationsClientWithBaseURI(sdkApi sdkEnv.Api) (*HuntRelationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "huntrelations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HuntRelationsClient: %+v", err)
	}

	return &HuntRelationsClient{
		Client: client,
	}, nil
}
