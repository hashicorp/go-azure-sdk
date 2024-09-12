package joinedteamphoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPhotoClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPhotoClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPhotoClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPhotoClient: %+v", err)
	}

	return &JoinedTeamPhotoClient{
		Client: client,
	}, nil
}
