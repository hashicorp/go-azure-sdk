package v2019_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/managedapplications/2019-07-01/applicationdefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedapplications/2019-07-01/applications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/managedapplications/2019-07-01/jitrequests"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ApplicationDefinitions *applicationdefinitions.ApplicationDefinitionsClient
	Applications           *applications.ApplicationsClient
	JitRequests            *jitrequests.JitRequestsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	applicationDefinitionsClient, err := applicationdefinitions.NewApplicationDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationDefinitions client: %+v", err)
	}
	configureFunc(applicationDefinitionsClient.Client)

	applicationsClient, err := applications.NewApplicationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Applications client: %+v", err)
	}
	configureFunc(applicationsClient.Client)

	jitRequestsClient, err := jitrequests.NewJitRequestsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JitRequests client: %+v", err)
	}
	configureFunc(jitRequestsClient.Client)

	return &Client{
		ApplicationDefinitions: applicationDefinitionsClient,
		Applications:           applicationsClient,
		JitRequests:            jitRequestsClient,
	}, nil
}
