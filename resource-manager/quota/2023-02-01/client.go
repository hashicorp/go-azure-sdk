package v2023_02_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2023-02-01/quotainformation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2023-02-01/quotarequests"
	"github.com/hashicorp/go-azure-sdk/resource-manager/quota/2023-02-01/usagesinformation"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	QuotaInformation  *quotainformation.QuotaInformationClient
	QuotaRequests     *quotarequests.QuotaRequestsClient
	UsagesInformation *usagesinformation.UsagesInformationClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	quotaInformationClient, err := quotainformation.NewQuotaInformationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QuotaInformation client: %+v", err)
	}
	configureFunc(quotaInformationClient.Client)

	quotaRequestsClient, err := quotarequests.NewQuotaRequestsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QuotaRequests client: %+v", err)
	}
	configureFunc(quotaRequestsClient.Client)

	usagesInformationClient, err := usagesinformation.NewUsagesInformationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UsagesInformation client: %+v", err)
	}
	configureFunc(usagesInformationClient.Client)

	return &Client{
		QuotaInformation:  quotaInformationClient,
		QuotaRequests:     quotaRequestsClient,
		UsagesInformation: usagesInformationClient,
	}, nil
}
