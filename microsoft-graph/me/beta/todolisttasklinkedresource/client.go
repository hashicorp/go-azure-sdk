package todolisttasklinkedresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoListTaskLinkedResourceClient struct {
	Client *msgraph.Client
}

func NewTodoListTaskLinkedResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*TodoListTaskLinkedResourceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "todolisttasklinkedresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TodoListTaskLinkedResourceClient: %+v", err)
	}

	return &TodoListTaskLinkedResourceClient{
		Client: client,
	}, nil
}
