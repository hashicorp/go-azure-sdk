package groupteamprimarychanneltab

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamPrimaryChannelTabClient struct {
	Client *msgraph.Client
}

func NewGroupTeamPrimaryChannelTabClientWithBaseURI(api sdkEnv.Api) (*GroupTeamPrimaryChannelTabClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamprimarychanneltab", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamPrimaryChannelTabClient: %+v", err)
	}

	return &GroupTeamPrimaryChannelTabClient{
		Client: client,
	}, nil
}
