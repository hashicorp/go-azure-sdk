package todolisttaskattachmentsessioncontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoListTaskAttachmentSessionContentClient struct {
	Client *msgraph.Client
}

func NewTodoListTaskAttachmentSessionContentClientWithBaseURI(sdkApi sdkEnv.Api) (*TodoListTaskAttachmentSessionContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "todolisttaskattachmentsessioncontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TodoListTaskAttachmentSessionContentClient: %+v", err)
	}

	return &TodoListTaskAttachmentSessionContentClient{
		Client: client,
	}, nil
}
