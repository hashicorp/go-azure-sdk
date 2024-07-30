package todo

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoClient struct {
	Client *msgraph.Client
}

func NewTodoClientWithBaseURI(sdkApi sdkEnv.Api) (*TodoClient, error) {
	client, err := msgraph.NewClient(sdkApi, "todo", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TodoClient: %+v", err)
	}

	return &TodoClient{
		Client: client,
	}, nil
}
