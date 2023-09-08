package userjoinedteamprimarychannelfilesfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPrimaryChannelFilesFolderClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPrimaryChannelFilesFolderClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPrimaryChannelFilesFolderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamprimarychannelfilesfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPrimaryChannelFilesFolderClient: %+v", err)
	}

	return &UserJoinedTeamPrimaryChannelFilesFolderClient{
		Client: client,
	}, nil
}
