package meonenote

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenote", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteClient: %+v", err)
	}

	return &MeOnenoteClient{
		Client: client,
	}, nil
}
