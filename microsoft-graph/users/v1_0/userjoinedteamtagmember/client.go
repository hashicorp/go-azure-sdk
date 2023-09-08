package userjoinedteamtagmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamTagMemberClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamTagMemberClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamTagMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamtagmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamTagMemberClient: %+v", err)
	}

	return &UserJoinedTeamTagMemberClient{
		Client: client,
	}, nil
}
