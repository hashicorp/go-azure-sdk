package teamtemplatedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamTemplateDefinitionClient struct {
	Client *msgraph.Client
}

func NewTeamTemplateDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamTemplateDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamtemplatedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamTemplateDefinitionClient: %+v", err)
	}

	return &TeamTemplateDefinitionClient{
		Client: client,
	}, nil
}
