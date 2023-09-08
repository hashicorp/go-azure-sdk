package userscopedrolememberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserScopedRoleMemberOfClient struct {
	Client *msgraph.Client
}

func NewUserScopedRoleMemberOfClientWithBaseURI(api sdkEnv.Api) (*UserScopedRoleMemberOfClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userscopedrolememberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserScopedRoleMemberOfClient: %+v", err)
	}

	return &UserScopedRoleMemberOfClient{
		Client: client,
	}, nil
}
