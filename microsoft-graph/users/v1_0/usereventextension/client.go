package usereventextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventExtensionClient struct {
	Client *msgraph.Client
}

func NewUserEventExtensionClientWithBaseURI(api sdkEnv.Api) (*UserEventExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventExtensionClient: %+v", err)
	}

	return &UserEventExtensionClient{
		Client: client,
	}, nil
}
