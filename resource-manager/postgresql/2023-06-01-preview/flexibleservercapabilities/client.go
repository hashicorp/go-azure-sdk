package flexibleservercapabilities

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FlexibleServerCapabilitiesClient struct {
	Client *resourcemanager.Client
}

func NewFlexibleServerCapabilitiesClientWithBaseURI(sdkApi sdkEnv.Api) (*FlexibleServerCapabilitiesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "flexibleservercapabilities", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FlexibleServerCapabilitiesClient: %+v", err)
	}

	return &FlexibleServerCapabilitiesClient{
		Client: client,
	}, nil
}
