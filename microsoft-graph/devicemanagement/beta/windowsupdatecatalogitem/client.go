package windowsupdatecatalogitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateCatalogItemClient struct {
	Client *msgraph.Client
}

func NewWindowsUpdateCatalogItemClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsUpdateCatalogItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsupdatecatalogitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsUpdateCatalogItemClient: %+v", err)
	}

	return &WindowsUpdateCatalogItemClient{
		Client: client,
	}, nil
}
