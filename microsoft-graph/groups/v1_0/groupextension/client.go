package groupextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupExtensionClient: %+v", err)
	}

	return &GroupExtensionClient{
		Client: client,
	}, nil
}
