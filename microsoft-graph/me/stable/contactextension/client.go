package contactextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactExtensionClient struct {
	Client *msgraph.Client
}

func NewContactExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*ContactExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "contactextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContactExtensionClient: %+v", err)
	}

	return &ContactExtensionClient{
		Client: client,
	}, nil
}
