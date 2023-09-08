package userprofileinterest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfileInterestClient struct {
	Client *msgraph.Client
}

func NewUserProfileInterestClientWithBaseURI(api sdkEnv.Api) (*UserProfileInterestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofileinterest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfileInterestClient: %+v", err)
	}

	return &UserProfileInterestClient{
		Client: client,
	}, nil
}
