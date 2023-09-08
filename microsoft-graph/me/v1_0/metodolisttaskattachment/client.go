package metodolisttaskattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTodoListTaskAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeTodoListTaskAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeTodoListTaskAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "metodolisttaskattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTodoListTaskAttachmentClient: %+v", err)
	}

	return &MeTodoListTaskAttachmentClient{
		Client: client,
	}, nil
}
