package mejoinedteamprimarychannelmessagehostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamPrimaryChannelMessageHostedContentClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamPrimaryChannelMessageHostedContentClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamPrimaryChannelMessageHostedContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamprimarychannelmessagehostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamPrimaryChannelMessageHostedContentClient: %+v", err)
	}

	return &MeJoinedTeamPrimaryChannelMessageHostedContentClient{
		Client: client,
	}, nil
}
