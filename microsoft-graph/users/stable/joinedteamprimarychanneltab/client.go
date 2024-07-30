package joinedteamprimarychanneltab

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelTabClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelTabClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelTabClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychanneltab", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelTabClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelTabClient{
		Client: client,
	}, nil
}
