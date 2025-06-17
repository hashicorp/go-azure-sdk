package defenderroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewDefenderRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*DefenderRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "defenderroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DefenderRoleDefinitionClient: %+v", err)
	}

	return &DefenderRoleDefinitionClient{
		Client: client,
	}, nil
}
