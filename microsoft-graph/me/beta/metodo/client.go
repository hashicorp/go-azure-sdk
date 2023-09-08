package metodo

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTodoClient struct {
	Client *msgraph.Client
}

func NewMeTodoClientWithBaseURI(api sdkEnv.Api) (*MeTodoClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "metodo", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTodoClient: %+v", err)
	}

	return &MeTodoClient{
		Client: client,
	}, nil
}
