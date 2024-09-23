package joinedteamprimarychannelfilesfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamPrimaryChannelFilesFolderClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamPrimaryChannelFilesFolderClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamPrimaryChannelFilesFolderClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamprimarychannelfilesfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamPrimaryChannelFilesFolderClient: %+v", err)
	}

	return &JoinedTeamPrimaryChannelFilesFolderClient{
		Client: client,
	}, nil
}
