package groupteamprimarychannelmessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamPrimaryChannelMessageClient struct {
	Client *msgraph.Client
}

func NewGroupTeamPrimaryChannelMessageClientWithBaseURI(api sdkEnv.Api) (*GroupTeamPrimaryChannelMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamprimarychannelmessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamPrimaryChannelMessageClient: %+v", err)
	}

	return &GroupTeamPrimaryChannelMessageClient{
		Client: client,
	}, nil
}
