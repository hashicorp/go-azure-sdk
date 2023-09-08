package userjoinedteamchannelfilesfoldercontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamChannelFilesFolderContentClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamChannelFilesFolderContentClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamChannelFilesFolderContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamchannelfilesfoldercontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamChannelFilesFolderContentClient: %+v", err)
	}

	return &UserJoinedTeamChannelFilesFolderContentClient{
		Client: client,
	}, nil
}
