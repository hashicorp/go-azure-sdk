package usermailfoldermessageextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderMessageExtensionClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderMessageExtensionClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderMessageExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfoldermessageextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderMessageExtensionClient: %+v", err)
	}

	return &UserMailFolderMessageExtensionClient{
		Client: client,
	}, nil
}
