package userprofilewebsite

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfileWebsiteClient struct {
	Client *msgraph.Client
}

func NewUserProfileWebsiteClientWithBaseURI(api sdkEnv.Api) (*UserProfileWebsiteClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofilewebsite", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfileWebsiteClient: %+v", err)
	}

	return &UserProfileWebsiteClient{
		Client: client,
	}, nil
}
