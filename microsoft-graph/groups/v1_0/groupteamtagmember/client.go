package groupteamtagmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamTagMemberClient struct {
	Client *msgraph.Client
}

func NewGroupTeamTagMemberClientWithBaseURI(api sdkEnv.Api) (*GroupTeamTagMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamtagmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamTagMemberClient: %+v", err)
	}

	return &GroupTeamTagMemberClient{
		Client: client,
	}, nil
}
