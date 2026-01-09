package todolisttask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoListTaskClient struct {
	Client *msgraph.Client
}

func NewTodoListTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*TodoListTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "todolisttask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TodoListTaskClient: %+v", err)
	}

	return &TodoListTaskClient{
		Client: client,
	}, nil
}
