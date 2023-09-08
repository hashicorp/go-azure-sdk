package memailfoldermessageextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMailFolderMessageExtensionClient struct {
	Client *msgraph.Client
}

func NewMeMailFolderMessageExtensionClientWithBaseURI(api sdkEnv.Api) (*MeMailFolderMessageExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memailfoldermessageextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMailFolderMessageExtensionClient: %+v", err)
	}

	return &MeMailFolderMessageExtensionClient{
		Client: client,
	}, nil
}
