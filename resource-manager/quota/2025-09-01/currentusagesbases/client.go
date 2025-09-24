package currentusagesbases

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CurrentUsagesBasesClient struct {
	Client *resourcemanager.Client
}

func NewCurrentUsagesBasesClientWithBaseURI(sdkApi sdkEnv.Api) (*CurrentUsagesBasesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "currentusagesbases", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CurrentUsagesBasesClient: %+v", err)
	}

	return &CurrentUsagesBasesClient{
		Client: client,
	}, nil
}
