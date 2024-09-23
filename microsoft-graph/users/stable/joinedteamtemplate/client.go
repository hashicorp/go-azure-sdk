package joinedteamtemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamTemplateClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamTemplateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamtemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamTemplateClient: %+v", err)
	}

	return &JoinedTeamTemplateClient{
		Client: client,
	}, nil
}
