package todolisttaskchecklistitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoListTaskChecklistItemClient struct {
	Client *msgraph.Client
}

func NewTodoListTaskChecklistItemClientWithBaseURI(sdkApi sdkEnv.Api) (*TodoListTaskChecklistItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "todolisttaskchecklistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TodoListTaskChecklistItemClient: %+v", err)
	}

	return &TodoListTaskChecklistItemClient{
		Client: client,
	}, nil
}
