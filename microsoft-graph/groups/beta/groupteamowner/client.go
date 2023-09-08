package groupteamowner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamOwnerClient struct {
	Client *msgraph.Client
}

func NewGroupTeamOwnerClientWithBaseURI(api sdkEnv.Api) (*GroupTeamOwnerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamowner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamOwnerClient: %+v", err)
	}

	return &GroupTeamOwnerClient{
		Client: client,
	}, nil
}
