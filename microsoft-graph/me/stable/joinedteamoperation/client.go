package joinedteamoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamOperationClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamOperationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamOperationClient: %+v", err)
	}

	return &JoinedTeamOperationClient{
		Client: client,
	}, nil
}
