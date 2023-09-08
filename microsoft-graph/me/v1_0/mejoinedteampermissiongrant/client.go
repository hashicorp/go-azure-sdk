package mejoinedteampermissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamPermissionGrantClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamPermissionGrantClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamPermissionGrantClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteampermissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamPermissionGrantClient: %+v", err)
	}

	return &MeJoinedTeamPermissionGrantClient{
		Client: client,
	}, nil
}
