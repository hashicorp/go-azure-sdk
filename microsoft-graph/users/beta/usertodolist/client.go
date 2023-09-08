package usertodolist

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTodoListClient struct {
	Client *msgraph.Client
}

func NewUserTodoListClientWithBaseURI(api sdkEnv.Api) (*UserTodoListClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usertodolist", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTodoListClient: %+v", err)
	}

	return &UserTodoListClient{
		Client: client,
	}, nil
}
