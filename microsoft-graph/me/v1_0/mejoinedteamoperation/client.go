package mejoinedteamoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamOperationClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamOperationClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamOperationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamOperationClient: %+v", err)
	}

	return &MeJoinedTeamOperationClient{
		Client: client,
	}, nil
}
