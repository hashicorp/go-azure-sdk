package joinedteampermissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPermissionGrantClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPermissionGrantClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPermissionGrantClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteampermissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPermissionGrantClient: %+v", err)
	}

	return &JoinedTeamPermissionGrantClient{
		Client: client,
	}, nil
}
