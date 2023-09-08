package userjoinedteamprimarychannelfilesfoldercontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedTeamPrimaryChannelFilesFolderContentClient struct {
	Client *msgraph.Client
}

func NewUserJoinedTeamPrimaryChannelFilesFolderContentClientWithBaseURI(api sdkEnv.Api) (*UserJoinedTeamPrimaryChannelFilesFolderContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedteamprimarychannelfilesfoldercontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedTeamPrimaryChannelFilesFolderContentClient: %+v", err)
	}

	return &UserJoinedTeamPrimaryChannelFilesFolderContentClient{
		Client: client,
	}, nil
}
