package invoices

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvoicesClient struct {
	Client *resourcemanager.Client
}

func NewInvoicesClientWithBaseURI(sdkApi sdkEnv.Api) (*InvoicesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "invoices", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvoicesClient: %+v", err)
	}

	return &InvoicesClient{
		Client: client,
	}, nil
}
