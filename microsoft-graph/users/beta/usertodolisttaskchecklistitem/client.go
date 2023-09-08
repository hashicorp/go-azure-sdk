package usertodolisttaskchecklistitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTodoListTaskChecklistItemClient struct {
	Client *msgraph.Client
}

func NewUserTodoListTaskChecklistItemClientWithBaseURI(api sdkEnv.Api) (*UserTodoListTaskChecklistItemClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usertodolisttaskchecklistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTodoListTaskChecklistItemClient: %+v", err)
	}

	return &UserTodoListTaskChecklistItemClient{
		Client: client,
	}, nil
}
