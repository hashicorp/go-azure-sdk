package teamprimarychannelfilesfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamPrimaryChannelFilesFolderClient struct {
	Client *msgraph.Client
}

func NewTeamPrimaryChannelFilesFolderClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamPrimaryChannelFilesFolderClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamprimarychannelfilesfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamPrimaryChannelFilesFolderClient: %+v", err)
	}

	return &TeamPrimaryChannelFilesFolderClient{
		Client: client,
	}, nil
}
