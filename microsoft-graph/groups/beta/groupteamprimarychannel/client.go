package groupteamprimarychannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamPrimaryChannelClient struct {
	Client *msgraph.Client
}

func NewGroupTeamPrimaryChannelClientWithBaseURI(api sdkEnv.Api) (*GroupTeamPrimaryChannelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamprimarychannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamPrimaryChannelClient: %+v", err)
	}

	return &GroupTeamPrimaryChannelClient{
		Client: client,
	}, nil
}
