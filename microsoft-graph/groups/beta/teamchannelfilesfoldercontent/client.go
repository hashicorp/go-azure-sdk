package teamchannelfilesfoldercontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamChannelFilesFolderContentClient struct {
	Client *msgraph.Client
}

func NewTeamChannelFilesFolderContentClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamChannelFilesFolderContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamchannelfilesfoldercontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamChannelFilesFolderContentClient: %+v", err)
	}

	return &TeamChannelFilesFolderContentClient{
		Client: client,
	}, nil
}
