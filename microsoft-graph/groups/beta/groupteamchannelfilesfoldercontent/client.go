package groupteamchannelfilesfoldercontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamChannelFilesFolderContentClient struct {
	Client *msgraph.Client
}

func NewGroupTeamChannelFilesFolderContentClientWithBaseURI(api sdkEnv.Api) (*GroupTeamChannelFilesFolderContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamchannelfilesfoldercontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamChannelFilesFolderContentClient: %+v", err)
	}

	return &GroupTeamChannelFilesFolderContentClient{
		Client: client,
	}, nil
}
