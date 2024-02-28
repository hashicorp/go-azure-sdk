package applicationwhitelistings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationWhitelistingsClient struct {
	Client *resourcemanager.Client
}

func NewApplicationWhitelistingsClientWithBaseURI(sdkApi sdkEnv.Api) (*ApplicationWhitelistingsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "applicationwhitelistings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationWhitelistingsClient: %+v", err)
	}

	return &ApplicationWhitelistingsClient{
		Client: client,
	}, nil
}
