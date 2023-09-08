package grouponenote

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOnenoteClient struct {
	Client *msgraph.Client
}

func NewGroupOnenoteClientWithBaseURI(api sdkEnv.Api) (*GroupOnenoteClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouponenote", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOnenoteClient: %+v", err)
	}

	return &GroupOnenoteClient{
		Client: client,
	}, nil
}
