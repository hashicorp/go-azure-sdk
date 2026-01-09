package contactfoldercontactextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactFolderContactExtensionClient struct {
	Client *msgraph.Client
}

func NewContactFolderContactExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*ContactFolderContactExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "contactfoldercontactextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContactFolderContactExtensionClient: %+v", err)
	}

	return &ContactFolderContactExtensionClient{
		Client: client,
	}, nil
}
