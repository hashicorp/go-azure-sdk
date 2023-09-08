package groupteamtemplatedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamTemplateDefinitionClient struct {
	Client *msgraph.Client
}

func NewGroupTeamTemplateDefinitionClientWithBaseURI(api sdkEnv.Api) (*GroupTeamTemplateDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamtemplatedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamTemplateDefinitionClient: %+v", err)
	}

	return &GroupTeamTemplateDefinitionClient{
		Client: client,
	}, nil
}
