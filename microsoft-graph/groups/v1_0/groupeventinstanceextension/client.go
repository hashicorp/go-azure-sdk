package groupeventinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupEventInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupEventInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventInstanceExtensionClient: %+v", err)
	}

	return &GroupEventInstanceExtensionClient{
		Client: client,
	}, nil
}
