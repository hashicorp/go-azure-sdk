package mejoinedteamphoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamPhotoClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamPhotoClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamPhotoClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamPhotoClient: %+v", err)
	}

	return &MeJoinedTeamPhotoClient{
		Client: client,
	}, nil
}
