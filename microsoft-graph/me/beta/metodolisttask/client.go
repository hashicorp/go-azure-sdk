package metodolisttask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTodoListTaskClient struct {
	Client *msgraph.Client
}

func NewMeTodoListTaskClientWithBaseURI(api sdkEnv.Api) (*MeTodoListTaskClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "metodolisttask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTodoListTaskClient: %+v", err)
	}

	return &MeTodoListTaskClient{
		Client: client,
	}, nil
}
