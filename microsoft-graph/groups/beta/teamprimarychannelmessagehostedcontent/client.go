package teamprimarychannelmessagehostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelMessageHostedContentClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelMessageHostedContentClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelMessageHostedContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelmessagehostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelMessageHostedContentClient: %+v", err)
	}

	return &TeamPrimaryChannelMessageHostedContentClient{
		Client: client,
	}, nil
}
