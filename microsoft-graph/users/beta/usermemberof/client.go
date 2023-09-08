package usermemberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMemberOfClient struct {
	Client *msgraph.Client
}

func NewUserMemberOfClientWithBaseURI(api sdkEnv.Api) (*UserMemberOfClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermemberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMemberOfClient: %+v", err)
	}

	return &UserMemberOfClient{
		Client: client,
	}, nil
}
