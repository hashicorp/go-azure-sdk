package mejoinedteamtemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamTemplateClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamTemplateClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamTemplateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamtemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamTemplateClient: %+v", err)
	}

	return &MeJoinedTeamTemplateClient{
		Client: client,
	}, nil
}
