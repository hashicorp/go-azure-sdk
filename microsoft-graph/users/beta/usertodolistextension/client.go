package usertodolistextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTodoListExtensionClient struct {
	Client *msgraph.Client
}

func NewUserTodoListExtensionClientWithBaseURI(api sdkEnv.Api) (*UserTodoListExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usertodolistextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTodoListExtensionClient: %+v", err)
	}

	return &UserTodoListExtensionClient{
		Client: client,
	}, nil
}
