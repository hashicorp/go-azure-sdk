package groupteamprimarychannelsharedwithteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamPrimaryChannelSharedWithTeamClient struct {
	Client *msgraph.Client
}

func NewGroupTeamPrimaryChannelSharedWithTeamClientWithBaseURI(api sdkEnv.Api) (*GroupTeamPrimaryChannelSharedWithTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamprimarychannelsharedwithteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamPrimaryChannelSharedWithTeamClient: %+v", err)
	}

	return &GroupTeamPrimaryChannelSharedWithTeamClient{
		Client: client,
	}, nil
}
