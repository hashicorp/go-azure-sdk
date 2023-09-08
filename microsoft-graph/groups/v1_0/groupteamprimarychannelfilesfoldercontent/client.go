package groupteamprimarychannelfilesfoldercontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamPrimaryChannelFilesFolderContentClient struct {
	Client *msgraph.Client
}

func NewGroupTeamPrimaryChannelFilesFolderContentClientWithBaseURI(api sdkEnv.Api) (*GroupTeamPrimaryChannelFilesFolderContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamprimarychannelfilesfoldercontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamPrimaryChannelFilesFolderContentClient: %+v", err)
	}

	return &GroupTeamPrimaryChannelFilesFolderContentClient{
		Client: client,
	}, nil
}
