package usagesinformation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsagesInformationClient struct {
	Client *resourcemanager.Client
}

func NewUsagesInformationClientWithBaseURI(sdkApi sdkEnv.Api) (*UsagesInformationClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "usagesinformation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UsagesInformationClient: %+v", err)
	}

	return &UsagesInformationClient{
		Client: client,
	}, nil
}
