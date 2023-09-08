package groupteamprimarychannelmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamPrimaryChannelMemberClient struct {
	Client *msgraph.Client
}

func NewGroupTeamPrimaryChannelMemberClientWithBaseURI(api sdkEnv.Api) (*GroupTeamPrimaryChannelMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamprimarychannelmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamPrimaryChannelMemberClient: %+v", err)
	}

	return &GroupTeamPrimaryChannelMemberClient{
		Client: client,
	}, nil
}
