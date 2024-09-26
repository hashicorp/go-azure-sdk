package securitystandards

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityStandardsClient struct {
	Client *resourcemanager.Client
}

func NewSecurityStandardsClientWithBaseURI(sdkApi sdkEnv.Api) (*SecurityStandardsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "securitystandards", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SecurityStandardsClient: %+v", err)
	}

	return &SecurityStandardsClient{
		Client: client,
	}, nil
}
