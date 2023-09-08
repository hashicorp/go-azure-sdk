package userprofileaccount

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfileAccountClient struct {
	Client *msgraph.Client
}

func NewUserProfileAccountClientWithBaseURI(api sdkEnv.Api) (*UserProfileAccountClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofileaccount", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfileAccountClient: %+v", err)
	}

	return &UserProfileAccountClient{
		Client: client,
	}, nil
}
