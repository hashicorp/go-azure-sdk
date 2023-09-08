package userpeople

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPeopleClient struct {
	Client *msgraph.Client
}

func NewUserPeopleClientWithBaseURI(api sdkEnv.Api) (*UserPeopleClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpeople", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPeopleClient: %+v", err)
	}

	return &UserPeopleClient{
		Client: client,
	}, nil
}
