package mecontactfoldercontactextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeContactFolderContactExtensionClient struct {
	Client *msgraph.Client
}

func NewMeContactFolderContactExtensionClientWithBaseURI(api sdkEnv.Api) (*MeContactFolderContactExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecontactfoldercontactextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeContactFolderContactExtensionClient: %+v", err)
	}

	return &MeContactFolderContactExtensionClient{
		Client: client,
	}, nil
}
