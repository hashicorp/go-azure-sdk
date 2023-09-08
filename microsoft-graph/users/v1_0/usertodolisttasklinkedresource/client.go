package usertodolisttasklinkedresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTodoListTaskLinkedResourceClient struct {
	Client *msgraph.Client
}

func NewUserTodoListTaskLinkedResourceClientWithBaseURI(api sdkEnv.Api) (*UserTodoListTaskLinkedResourceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usertodolisttasklinkedresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTodoListTaskLinkedResourceClient: %+v", err)
	}

	return &UserTodoListTaskLinkedResourceClient{
		Client: client,
	}, nil
}
