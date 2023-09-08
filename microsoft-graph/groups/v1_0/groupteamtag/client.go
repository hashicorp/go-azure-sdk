package groupteamtag

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamTagClient struct {
	Client *msgraph.Client
}

func NewGroupTeamTagClientWithBaseURI(api sdkEnv.Api) (*GroupTeamTagClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamtag", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamTagClient: %+v", err)
	}

	return &GroupTeamTagClient{
		Client: client,
	}, nil
}
