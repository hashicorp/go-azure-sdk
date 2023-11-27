package timezones

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeZonesClient struct {
	Client *resourcemanager.Client
}

func NewTimeZonesClientWithBaseURI(sdkApi sdkEnv.Api) (*TimeZonesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "timezones", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TimeZonesClient: %+v", err)
	}

	return &TimeZonesClient{
		Client: client,
	}, nil
}
