package usereventinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserEventInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserEventInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventInstanceExtensionClient: %+v", err)
	}

	return &UserEventInstanceExtensionClient{
		Client: client,
	}, nil
}
