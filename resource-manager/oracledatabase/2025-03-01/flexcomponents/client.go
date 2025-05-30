package flexcomponents

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FlexComponentsClient struct {
	Client *resourcemanager.Client
}

func NewFlexComponentsClientWithBaseURI(sdkApi sdkEnv.Api) (*FlexComponentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "flexcomponents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FlexComponentsClient: %+v", err)
	}

	return &FlexComponentsClient{
		Client: client,
	}, nil
}
