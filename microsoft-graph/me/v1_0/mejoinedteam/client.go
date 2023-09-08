package mejoinedteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamClient: %+v", err)
	}

	return &MeJoinedTeamClient{
		Client: client,
	}, nil
}
