package joinedteamprimarychannelmessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelMessageClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelMessageClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelMessageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychannelmessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelMessageClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelMessageClient{
		Client: client,
	}, nil
}
