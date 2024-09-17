package templatespecs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateSpecsClient struct {
	Client *resourcemanager.Client
}

func NewTemplateSpecsClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateSpecsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "templatespecs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateSpecsClient: %+v", err)
	}

	return &TemplateSpecsClient{
		Client: client,
	}, nil
}
