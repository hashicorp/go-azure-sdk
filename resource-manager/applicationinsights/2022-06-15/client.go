package v2022_06_15

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2022-06-15/webtestsapis"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	WebTestsAPIs *webtestsapis.WebTestsAPIsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	webTestsAPIsClient := webtestsapis.NewWebTestsAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&webTestsAPIsClient.Client)

	return Client{
		WebTestsAPIs: &webTestsAPIsClient,
	}
}
