package teamprimarychannelfilesfoldercontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelFilesFolderContentStreamClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelFilesFolderContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelFilesFolderContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelfilesfoldercontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelFilesFolderContentStreamClient: %+v", err)
	}

	return &TeamPrimaryChannelFilesFolderContentStreamClient{
		Client: client,
	}, nil
}
