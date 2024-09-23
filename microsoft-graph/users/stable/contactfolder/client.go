package contactfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactFolderClient struct {
	Client *msgraph.Client
}

func NewContactFolderClientWithBaseURI(sdkApi sdkEnv.Api) (*ContactFolderClient, error) {
	client, err := msgraph.NewClient(sdkApi, "contactfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContactFolderClient: %+v", err)
	}

	return &ContactFolderClient{
		Client: client,
	}, nil
}
