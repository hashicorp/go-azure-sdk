package meapproleassignedresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAppRoleAssignedResourceClient struct {
	Client *msgraph.Client
}

func NewMeAppRoleAssignedResourceClientWithBaseURI(api sdkEnv.Api) (*MeAppRoleAssignedResourceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meapproleassignedresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAppRoleAssignedResourceClient: %+v", err)
	}

	return &MeAppRoleAssignedResourceClient{
		Client: client,
	}, nil
}
