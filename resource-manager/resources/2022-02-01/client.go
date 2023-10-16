package v2022_02_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-02-01/templatespecs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-02-01/templatespecversions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	TemplateSpecVersions *templatespecversions.TemplateSpecVersionsClient
	TemplateSpecs        *templatespecs.TemplateSpecsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	templateSpecVersionsClient, err := templatespecversions.NewTemplateSpecVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateSpecVersions client: %+v", err)
	}
	configureFunc(templateSpecVersionsClient.Client)

	templateSpecsClient, err := templatespecs.NewTemplateSpecsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TemplateSpecs client: %+v", err)
	}
	configureFunc(templateSpecsClient.Client)

	return &Client{
		TemplateSpecVersions: templateSpecVersionsClient,
		TemplateSpecs:        templateSpecsClient,
	}, nil
}
