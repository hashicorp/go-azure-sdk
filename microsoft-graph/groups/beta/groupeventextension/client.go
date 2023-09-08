package groupeventextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupEventExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupEventExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventExtensionClient: %+v", err)
	}

	return &GroupEventExtensionClient{
		Client: client,
	}, nil
}
