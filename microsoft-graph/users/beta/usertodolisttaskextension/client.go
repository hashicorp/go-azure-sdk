package usertodolisttaskextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTodoListTaskExtensionClient struct {
	Client *msgraph.Client
}

func NewUserTodoListTaskExtensionClientWithBaseURI(api sdkEnv.Api) (*UserTodoListTaskExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usertodolisttaskextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTodoListTaskExtensionClient: %+v", err)
	}

	return &UserTodoListTaskExtensionClient{
		Client: client,
	}, nil
}
