package userjoinedteamprimarychannelmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPrimaryChannelMemberClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPrimaryChannelMemberClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPrimaryChannelMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamprimarychannelmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPrimaryChannelMemberClient: %+v", err)
	}

	return &UserJoinedTeamPrimaryChannelMemberClient{
		Client: client,
	}, nil
}
