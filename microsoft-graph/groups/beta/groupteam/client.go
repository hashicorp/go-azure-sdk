package groupteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamClient struct {
	Client *msgraph.Client
}

func NewGroupTeamClientWithBaseURI(api sdkEnv.Api) (*GroupTeamClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamClient: %+v", err)
	}

	return &GroupTeamClient{
		Client: client,
	}, nil
}
