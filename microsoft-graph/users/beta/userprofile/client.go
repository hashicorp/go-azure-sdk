package userprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfileClient struct {
	Client *msgraph.Client
}

func NewUserProfileClientWithBaseURI(api sdkEnv.Api) (*UserProfileClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfileClient: %+v", err)
	}

	return &UserProfileClient{
		Client: client,
	}, nil
}
