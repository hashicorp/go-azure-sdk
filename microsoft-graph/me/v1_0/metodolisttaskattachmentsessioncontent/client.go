package metodolisttaskattachmentsessioncontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTodoListTaskAttachmentSessionContentClient struct {
	Client *msgraph.Client
}

func NewMeTodoListTaskAttachmentSessionContentClientWithBaseURI(api sdkEnv.Api) (*MeTodoListTaskAttachmentSessionContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "metodolisttaskattachmentsessioncontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTodoListTaskAttachmentSessionContentClient: %+v", err)
	}

	return &MeTodoListTaskAttachmentSessionContentClient{
		Client: client,
	}, nil
}
