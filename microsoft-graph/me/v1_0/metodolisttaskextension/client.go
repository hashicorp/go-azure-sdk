package metodolisttaskextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTodoListTaskExtensionClient struct {
	Client *msgraph.Client
}

func NewMeTodoListTaskExtensionClientWithBaseURI(api sdkEnv.Api) (*MeTodoListTaskExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "metodolisttaskextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTodoListTaskExtensionClient: %+v", err)
	}

	return &MeTodoListTaskExtensionClient{
		Client: client,
	}, nil
}
