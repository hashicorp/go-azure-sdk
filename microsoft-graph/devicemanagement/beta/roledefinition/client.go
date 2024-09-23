package roledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*RoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "roledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleDefinitionClient: %+v", err)
	}

	return &RoleDefinitionClient{
		Client: client,
	}, nil
}
