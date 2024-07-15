package dotnetcomponents

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DotNetComponentsClient struct {
	Client *resourcemanager.Client
}

func NewDotNetComponentsClientWithBaseURI(sdkApi sdkEnv.Api) (*DotNetComponentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "dotnetcomponents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DotNetComponentsClient: %+v", err)
	}

	return &DotNetComponentsClient{
		Client: client,
	}, nil
}
