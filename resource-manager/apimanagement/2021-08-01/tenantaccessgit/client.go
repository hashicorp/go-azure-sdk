package tenantaccessgit

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantAccessGitClient struct {
	Client *resourcemanager.Client
}

func NewTenantAccessGitClientWithBaseURI(api environments.Api) (*TenantAccessGitClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "tenantaccessgit", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TenantAccessGitClient: %+v", err)
	}

	return &TenantAccessGitClient{
		Client: client,
	}, nil
}
