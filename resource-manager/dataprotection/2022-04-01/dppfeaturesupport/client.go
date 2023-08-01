package dppfeaturesupport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DppFeatureSupportClient struct {
	Client *resourcemanager.Client
}

func NewDppFeatureSupportClientWithBaseURI(sdkApi sdkEnv.Api) (*DppFeatureSupportClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "dppfeaturesupport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DppFeatureSupportClient: %+v", err)
	}

	return &DppFeatureSupportClient{
		Client: client,
	}, nil
}
