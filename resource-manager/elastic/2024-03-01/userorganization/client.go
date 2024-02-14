package userorganization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOrganizationClient struct {
	Client *resourcemanager.Client
}

func NewUserOrganizationClientWithBaseURI(sdkApi sdkEnv.Api) (*UserOrganizationClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "userorganization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOrganizationClient: %+v", err)
	}

	return &UserOrganizationClient{
		Client: client,
	}, nil
}
