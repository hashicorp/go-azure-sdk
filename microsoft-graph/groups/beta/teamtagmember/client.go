package teamtagmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamTagMemberClient struct {
	Client *msgraph.Client
}

func NewTeamTagMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamTagMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamtagmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamTagMemberClient: %+v", err)
	}

	return &TeamTagMemberClient{
		Client: client,
	}, nil
}
