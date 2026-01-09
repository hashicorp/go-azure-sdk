package teamprimarychanneltab

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelTabClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelTabClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelTabClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychanneltab", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelTabClient: %+v", err)
	}

	return &TeamPrimaryChannelTabClient{
		Client: client,
	}, nil
}
