package userprofileproject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfileProjectClient struct {
	Client *msgraph.Client
}

func NewUserProfileProjectClientWithBaseURI(api sdkEnv.Api) (*UserProfileProjectClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofileproject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfileProjectClient: %+v", err)
	}

	return &UserProfileProjectClient{
		Client: client,
	}, nil
}
