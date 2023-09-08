package usertodo

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTodoClient struct {
	Client *msgraph.Client
}

func NewUserTodoClientWithBaseURI(api sdkEnv.Api) (*UserTodoClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usertodo", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTodoClient: %+v", err)
	}

	return &UserTodoClient{
		Client: client,
	}, nil
}
