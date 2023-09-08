package metodolisttasklinkedresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTodoListTaskLinkedResourceClient struct {
	Client *msgraph.Client
}

func NewMeTodoListTaskLinkedResourceClientWithBaseURI(api sdkEnv.Api) (*MeTodoListTaskLinkedResourceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "metodolisttasklinkedresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTodoListTaskLinkedResourceClient: %+v", err)
	}

	return &MeTodoListTaskLinkedResourceClient{
		Client: client,
	}, nil
}
