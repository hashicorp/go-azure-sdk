package userjoinedteamgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamGroupClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamGroupClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamGroupClient: %+v", err)
	}

	return &UserJoinedTeamGroupClient{
		Client: client,
	}, nil
}
