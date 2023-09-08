package usercontactfoldercontactextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserContactFolderContactExtensionClient struct {
	Client *msgraph.Client
}

func NewUserContactFolderContactExtensionClientWithBaseURI(api sdkEnv.Api) (*UserContactFolderContactExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercontactfoldercontactextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserContactFolderContactExtensionClient: %+v", err)
	}

	return &UserContactFolderContactExtensionClient{
		Client: client,
	}, nil
}
