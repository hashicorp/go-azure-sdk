package todolisttaskextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoListTaskExtensionClient struct {
	Client *msgraph.Client
}

func NewTodoListTaskExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*TodoListTaskExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "todolisttaskextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TodoListTaskExtensionClient: %+v", err)
	}

	return &TodoListTaskExtensionClient{
		Client: client,
	}, nil
}
