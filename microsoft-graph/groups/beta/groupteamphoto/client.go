package groupteamphoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamPhotoClient struct {
	Client *msgraph.Client
}

func NewGroupTeamPhotoClientWithBaseURI(api sdkEnv.Api) (*GroupTeamPhotoClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamPhotoClient: %+v", err)
	}

	return &GroupTeamPhotoClient{
		Client: client,
	}, nil
}
