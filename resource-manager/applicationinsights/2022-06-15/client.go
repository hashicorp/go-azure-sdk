package v2022_06_15

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2022-06-15/webtestsapis"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	WebTestsAPIs *webtestsapis.WebTestsAPIsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	webTestsAPIsClient, err := webtestsapis.NewWebTestsAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WebTestsAPIs client: %+v", err)
	}
	configureFunc(webTestsAPIsClient.Client)

	return &Client{
		WebTestsAPIs: webTestsAPIsClient,
	}, nil
}
