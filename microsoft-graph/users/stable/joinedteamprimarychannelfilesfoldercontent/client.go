package joinedteamprimarychannelfilesfoldercontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelFilesFolderContentClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelFilesFolderContentClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelFilesFolderContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychannelfilesfoldercontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelFilesFolderContentClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelFilesFolderContentClient{
		Client: client,
	}, nil
}
