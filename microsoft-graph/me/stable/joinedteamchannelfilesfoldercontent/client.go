package joinedteamchannelfilesfoldercontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamChannelFilesFolderContentClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamChannelFilesFolderContentClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamChannelFilesFolderContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamchannelfilesfoldercontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamChannelFilesFolderContentClient: %+v", err)
	}

	return &JoinedTeamChannelFilesFolderContentClient{
		Client: client,
	}, nil
}
