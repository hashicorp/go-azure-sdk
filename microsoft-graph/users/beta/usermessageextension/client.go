package usermessageextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMessageExtensionClient struct {
	Client *msgraph.Client
}

func NewUserMessageExtensionClientWithBaseURI(api sdkEnv.Api) (*UserMessageExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermessageextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMessageExtensionClient: %+v", err)
	}

	return &UserMessageExtensionClient{
		Client: client,
	}, nil
}
