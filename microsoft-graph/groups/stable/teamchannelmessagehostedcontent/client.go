package teamchannelmessagehostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelMessageHostedContentClient struct {
	Client *msgraph.Client
}

func NewTeamChannelMessageHostedContentClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelMessageHostedContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelmessagehostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelMessageHostedContentClient: %+v", err)
	}

	return &TeamChannelMessageHostedContentClient{
		Client: client,
	}, nil
}
