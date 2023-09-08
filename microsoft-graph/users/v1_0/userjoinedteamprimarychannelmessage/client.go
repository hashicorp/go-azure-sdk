package userjoinedteamprimarychannelmessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPrimaryChannelMessageClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPrimaryChannelMessageClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPrimaryChannelMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamprimarychannelmessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPrimaryChannelMessageClient: %+v", err)
	}

	return &UserJoinedTeamPrimaryChannelMessageClient{
		Client: client,
	}, nil
}
