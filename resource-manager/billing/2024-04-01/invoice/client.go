package invoice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvoiceClient struct {
	Client *resourcemanager.Client
}

func NewInvoiceClientWithBaseURI(sdkApi sdkEnv.Api) (*InvoiceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "invoice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvoiceClient: %+v", err)
	}

	return &InvoiceClient{
		Client: client,
	}, nil
}
