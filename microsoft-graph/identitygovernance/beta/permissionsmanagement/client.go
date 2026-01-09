package permissionsmanagement

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsManagementClient struct {
	Client *msgraph.Client
}

func NewPermissionsManagementClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionsManagementClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissionsmanagement", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionsManagementClient: %+v", err)
	}

	return &PermissionsManagementClient{
		Client: client,
	}, nil
}
