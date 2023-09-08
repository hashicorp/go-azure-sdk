package userjoinedteampermissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPermissionGrantClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPermissionGrantClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPermissionGrantClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteampermissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPermissionGrantClient: %+v", err)
	}

	return &UserJoinedTeamPermissionGrantClient{
		Client: client,
	}, nil
}
