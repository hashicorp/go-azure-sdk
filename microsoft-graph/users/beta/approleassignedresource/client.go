package approleassignedresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppRoleAssignedResourceClient struct {
	Client *msgraph.Client
}

func NewAppRoleAssignedResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*AppRoleAssignedResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "approleassignedresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppRoleAssignedResourceClient: %+v", err)
	}

	return &AppRoleAssignedResourceClient{
		Client: client,
	}, nil
}
