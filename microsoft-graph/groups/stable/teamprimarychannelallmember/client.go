package teamprimarychannelallmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelAllMemberClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelAllMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelAllMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelallmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelAllMemberClient: %+v", err)
	}

	return &TeamPrimaryChannelAllMemberClient{
		Client: client,
	}, nil
}
