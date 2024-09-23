package joinedteamtagmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamTagMemberClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamTagMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamTagMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamtagmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamTagMemberClient: %+v", err)
	}

	return &JoinedTeamTagMemberClient{
		Client: client,
	}, nil
}
