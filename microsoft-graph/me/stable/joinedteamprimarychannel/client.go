package joinedteamprimarychannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelClient{
		Client: client,
	}, nil
}
