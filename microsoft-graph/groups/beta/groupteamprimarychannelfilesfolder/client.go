package groupteamprimarychannelfilesfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamPrimaryChannelFilesFolderClient struct {
	Client *msgraph.Client
}

func NewGroupTeamPrimaryChannelFilesFolderClientWithBaseURI(api sdkEnv.Api) (*GroupTeamPrimaryChannelFilesFolderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamprimarychannelfilesfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamPrimaryChannelFilesFolderClient: %+v", err)
	}

	return &GroupTeamPrimaryChannelFilesFolderClient{
		Client: client,
	}, nil
}
