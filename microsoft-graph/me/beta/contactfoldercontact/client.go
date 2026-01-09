package contactfoldercontact

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactFolderContactClient struct {
	Client *msgraph.Client
}

func NewContactFolderContactClientWithBaseURI(sdkApi sdkEnv.Api) (*ContactFolderContactClient, error) {
	client, err := msgraph.NewClient(sdkApi, "contactfoldercontact", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContactFolderContactClient: %+v", err)
	}

	return &ContactFolderContactClient{
		Client: client,
	}, nil
}
