package teamprimarychannelfilesfoldercontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelFilesFolderContentClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelFilesFolderContentClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelFilesFolderContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelfilesfoldercontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelFilesFolderContentClient: %+v", err)
	}

	return &TeamPrimaryChannelFilesFolderContentClient{
		Client: client,
	}, nil
}
