package userprofilepatent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfilePatentClient struct {
	Client *msgraph.Client
}

func NewUserProfilePatentClientWithBaseURI(api sdkEnv.Api) (*UserProfilePatentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofilepatent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfilePatentClient: %+v", err)
	}

	return &UserProfilePatentClient{
		Client: client,
	}, nil
}
