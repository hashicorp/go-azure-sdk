package metodolist

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTodoListClient struct {
	Client *msgraph.Client
}

func NewMeTodoListClientWithBaseURI(api sdkEnv.Api) (*MeTodoListClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "metodolist", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTodoListClient: %+v", err)
	}

	return &MeTodoListClient{
		Client: client,
	}, nil
}
