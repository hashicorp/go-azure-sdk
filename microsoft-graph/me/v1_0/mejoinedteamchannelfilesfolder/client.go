package mejoinedteamchannelfilesfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeJoinedTeamChannelFilesFolderClient struct {
	Client *msgraph.Client
}

func NewMeJoinedTeamChannelFilesFolderClientWithBaseURI(api sdkEnv.Api) (*MeJoinedTeamChannelFilesFolderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mejoinedteamchannelfilesfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeJoinedTeamChannelFilesFolderClient: %+v", err)
	}

	return &MeJoinedTeamChannelFilesFolderClient{
		Client: client,
	}, nil
}
