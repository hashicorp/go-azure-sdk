package groupteamprimarychanneltabteamsapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamPrimaryChannelTabTeamsAppClient struct {
	Client *msgraph.Client
}

func NewGroupTeamPrimaryChannelTabTeamsAppClientWithBaseURI(api sdkEnv.Api) (*GroupTeamPrimaryChannelTabTeamsAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamprimarychanneltabteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamPrimaryChannelTabTeamsAppClient: %+v", err)
	}

	return &GroupTeamPrimaryChannelTabTeamsAppClient{
		Client: client,
	}, nil
}
