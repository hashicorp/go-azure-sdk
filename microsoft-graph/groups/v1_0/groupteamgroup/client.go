package groupteamgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamGroupClient struct {
	Client *msgraph.Client
}

func NewGroupTeamGroupClientWithBaseURI(api sdkEnv.Api) (*GroupTeamGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamGroupClient: %+v", err)
	}

	return &GroupTeamGroupClient{
		Client: client,
	}, nil
}
