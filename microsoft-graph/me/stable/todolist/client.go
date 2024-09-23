package todolist

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoListClient struct {
	Client *msgraph.Client
}

func NewTodoListClientWithBaseURI(sdkApi sdkEnv.Api) (*TodoListClient, error) {
	client, err := msgraph.NewClient(sdkApi, "todolist", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TodoListClient: %+v", err)
	}

	return &TodoListClient{
		Client: client,
	}, nil
}
