package teamworkassociatedteam

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkAssociatedTeamClient struct {
	Client *msgraph.Client
}

func NewTeamworkAssociatedTeamClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamworkAssociatedTeamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamworkassociatedteam", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamworkAssociatedTeamClient: %+v", err)
	}

	return &TeamworkAssociatedTeamClient{
		Client: client,
	}, nil
}
