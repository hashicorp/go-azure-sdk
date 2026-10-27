package hcpopenshiftversions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HcpOpenShiftVersionsClient struct {
	Client *resourcemanager.Client
}

func NewHcpOpenShiftVersionsClientWithBaseURI(sdkApi sdkEnv.Api) (*HcpOpenShiftVersionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "hcpopenshiftversions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HcpOpenShiftVersionsClient: %+v", err)
	}

	return &HcpOpenShiftVersionsClient{
		Client: client,
	}, nil
}
