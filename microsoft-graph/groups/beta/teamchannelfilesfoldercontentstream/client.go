package teamchannelfilesfoldercontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelFilesFolderContentStreamClient struct {
	Client *msgraph.Client
}

func NewTeamChannelFilesFolderContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelFilesFolderContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelfilesfoldercontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelFilesFolderContentStreamClient: %+v", err)
	}

	return &TeamChannelFilesFolderContentStreamClient{
		Client: client,
	}, nil
}
