package meteamworkassociatedteamteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTeamworkAssociatedTeamTeamClient struct {
	Client *msgraph.Client
}

func NewMeTeamworkAssociatedTeamTeamClientWithBaseURI(api sdkEnv.Api) (*MeTeamworkAssociatedTeamTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meteamworkassociatedteamteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTeamworkAssociatedTeamTeamClient: %+v", err)
	}

	return &MeTeamworkAssociatedTeamTeamClient{
		Client: client,
	}, nil
}
