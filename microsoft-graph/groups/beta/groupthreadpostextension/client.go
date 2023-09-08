package groupthreadpostextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupThreadPostExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupThreadPostExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupThreadPostExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupthreadpostextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupThreadPostExtensionClient: %+v", err)
	}

	return &GroupThreadPostExtensionClient{
		Client: client,
	}, nil
}
