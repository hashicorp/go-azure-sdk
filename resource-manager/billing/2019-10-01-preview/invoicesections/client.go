package invoicesections

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvoiceSectionsClient struct {
	Client *resourcemanager.Client
}

func NewInvoiceSectionsClientWithBaseURI(sdkApi sdkEnv.Api) (*InvoiceSectionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "invoicesections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvoiceSectionsClient: %+v", err)
	}

	return &InvoiceSectionsClient{
		Client: client,
	}, nil
}
