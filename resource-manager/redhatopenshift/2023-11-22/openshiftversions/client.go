package openshiftversions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftVersionsClient struct {
	Client *resourcemanager.Client
}

func NewOpenShiftVersionsClientWithBaseURI(sdkApi sdkEnv.Api) (*OpenShiftVersionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "openshiftversions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OpenShiftVersionsClient: %+v", err)
	}

	return &OpenShiftVersionsClient{
		Client: client,
	}, nil
}
