package userjoinedteamtemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamTemplateClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamTemplateClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamTemplateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamtemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamTemplateClient: %+v", err)
	}

	return &UserJoinedTeamTemplateClient{
		Client: client,
	}, nil
}
