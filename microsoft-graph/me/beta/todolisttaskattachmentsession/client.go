package todolisttaskattachmentsession

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoListTaskAttachmentSessionClient struct {
	Client *msgraph.Client
}

func NewTodoListTaskAttachmentSessionClientWithBaseURI(sdkApi sdkEnv.Api) (*TodoListTaskAttachmentSessionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "todolisttaskattachmentsession", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TodoListTaskAttachmentSessionClient: %+v", err)
	}

	return &TodoListTaskAttachmentSessionClient{
		Client: client,
	}, nil
}
