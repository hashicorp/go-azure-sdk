package mhsmlistregions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MHSMListRegionsClient struct {
	Client *resourcemanager.Client
}

func NewMHSMListRegionsClientWithBaseURI(sdkApi sdkEnv.Api) (*MHSMListRegionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "mhsmlistregions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MHSMListRegionsClient: %+v", err)
	}

	return &MHSMListRegionsClient{
		Client: client,
	}, nil
}
