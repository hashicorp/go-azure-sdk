package usertodolisttaskattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTodoListTaskAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserTodoListTaskAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserTodoListTaskAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usertodolisttaskattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTodoListTaskAttachmentClient: %+v", err)
	}

	return &UserTodoListTaskAttachmentClient{
		Client: client,
	}, nil
}
