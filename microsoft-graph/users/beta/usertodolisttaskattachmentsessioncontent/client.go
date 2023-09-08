package usertodolisttaskattachmentsessioncontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTodoListTaskAttachmentSessionContentClient struct {
	Client *msgraph.Client
}

func NewUserTodoListTaskAttachmentSessionContentClientWithBaseURI(api sdkEnv.Api) (*UserTodoListTaskAttachmentSessionContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usertodolisttaskattachmentsessioncontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTodoListTaskAttachmentSessionContentClient: %+v", err)
	}

	return &UserTodoListTaskAttachmentSessionContentClient{
		Client: client,
	}, nil
}
