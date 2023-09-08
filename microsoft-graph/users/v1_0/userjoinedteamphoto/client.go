package userjoinedteamphoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPhotoClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPhotoClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPhotoClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPhotoClient: %+v", err)
	}

	return &UserJoinedTeamPhotoClient{
		Client: client,
	}, nil
}
