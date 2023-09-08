package metodolisttaskattachmentsession

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTodoListTaskAttachmentSessionClient struct {
	Client *msgraph.Client
}

func NewMeTodoListTaskAttachmentSessionClientWithBaseURI(api sdkEnv.Api) (*MeTodoListTaskAttachmentSessionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "metodolisttaskattachmentsession", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTodoListTaskAttachmentSessionClient: %+v", err)
	}

	return &MeTodoListTaskAttachmentSessionClient{
		Client: client,
	}, nil
}
