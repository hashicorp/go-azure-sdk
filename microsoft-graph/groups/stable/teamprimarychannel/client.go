package teamprimarychannel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelClient: %+v", err)
	}

	return &TeamPrimaryChannelClient{
		Client: client,
	}, nil
}
