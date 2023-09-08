package userjoinedgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserJoinedGroupClient struct {
	Client *msgraph.Client
}

func NewUserJoinedGroupClientWithBaseURI(api sdkEnv.Api) (*UserJoinedGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userjoinedgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserJoinedGroupClient: %+v", err)
	}

	return &UserJoinedGroupClient{
		Client: client,
	}, nil
}
