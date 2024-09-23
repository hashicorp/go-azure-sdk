package joinedteammember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamMemberClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteammember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamMemberClient: %+v", err)
	}

	return &JoinedTeamMemberClient{
		Client: client,
	}, nil
}
