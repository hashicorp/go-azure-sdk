package mejoinedteamtagmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamTagMemberClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamTagMemberClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamTagMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamtagmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamTagMemberClient: %+v", err)
	}

	return &MeJoinedTeamTagMemberClient{
		Client: client,
	}, nil
}
