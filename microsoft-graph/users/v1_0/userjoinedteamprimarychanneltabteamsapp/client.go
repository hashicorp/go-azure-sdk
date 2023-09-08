package userjoinedteamprimarychanneltabteamsapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPrimaryChannelTabTeamsAppClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPrimaryChannelTabTeamsAppClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPrimaryChannelTabTeamsAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamprimarychanneltabteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPrimaryChannelTabTeamsAppClient: %+v", err)
	}

	return &UserJoinedTeamPrimaryChannelTabTeamsAppClient{
		Client: client,
	}, nil
}
