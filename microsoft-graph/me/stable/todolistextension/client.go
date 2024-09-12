package todolistextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoListExtensionClient struct {
	Client *msgraph.Client
}

func NewTodoListExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*TodoListExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "todolistextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TodoListExtensionClient: %+v", err)
	}

	return &TodoListExtensionClient{
		Client: client,
	}, nil
}
