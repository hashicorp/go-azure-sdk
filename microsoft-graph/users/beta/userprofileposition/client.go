package userprofileposition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfilePositionClient struct {
	Client *msgraph.Client
}

func NewUserProfilePositionClientWithBaseURI(api sdkEnv.Api) (*UserProfilePositionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofileposition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfilePositionClient: %+v", err)
	}

	return &UserProfilePositionClient{
		Client: client,
	}, nil
}
