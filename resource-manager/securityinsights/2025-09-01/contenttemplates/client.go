package contenttemplates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentTemplatesClient struct {
	Client *resourcemanager.Client
}

func NewContentTemplatesClientWithBaseURI(sdkApi sdkEnv.Api) (*ContentTemplatesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "contenttemplates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContentTemplatesClient: %+v", err)
	}

	return &ContentTemplatesClient{
		Client: client,
	}, nil
}
