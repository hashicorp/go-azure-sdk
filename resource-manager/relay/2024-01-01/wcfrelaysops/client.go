package wcfrelaysops

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WcfRelaysOpsClient struct {
	Client *resourcemanager.Client
}

func NewWcfRelaysOpsClientWithBaseURI(sdkApi sdkEnv.Api) (*WcfRelaysOpsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "wcfrelaysops", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WcfRelaysOpsClient: %+v", err)
	}

	return &WcfRelaysOpsClient{
		Client: client,
	}, nil
}
