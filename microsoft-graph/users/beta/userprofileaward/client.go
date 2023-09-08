package userprofileaward

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfileAwardClient struct {
	Client *msgraph.Client
}

func NewUserProfileAwardClientWithBaseURI(api sdkEnv.Api) (*UserProfileAwardClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofileaward", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfileAwardClient: %+v", err)
	}

	return &UserProfileAwardClient{
		Client: client,
	}, nil
}
