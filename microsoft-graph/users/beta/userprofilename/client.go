package userprofilename

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfileNameClient struct {
	Client *msgraph.Client
}

func NewUserProfileNameClientWithBaseURI(api sdkEnv.Api) (*UserProfileNameClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofilename", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfileNameClient: %+v", err)
	}

	return &UserProfileNameClient{
		Client: client,
	}, nil
}
