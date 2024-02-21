package contentproducttemplates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentProductTemplatesClient struct {
	Client *resourcemanager.Client
}

func NewContentProductTemplatesClientWithBaseURI(sdkApi sdkEnv.Api) (*ContentProductTemplatesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "contentproducttemplates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContentProductTemplatesClient: %+v", err)
	}

	return &ContentProductTemplatesClient{
		Client: client,
	}, nil
}
