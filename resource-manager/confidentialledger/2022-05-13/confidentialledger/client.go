package confidentialledger

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfidentialLedgerClient struct {
	Client *resourcemanager.Client
}

func NewConfidentialLedgerClientWithBaseURI(api sdkEnv.Api) (*ConfidentialLedgerClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "confidentialledger", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConfidentialLedgerClient: %+v", err)
	}

	return &ConfidentialLedgerClient{
		Client: client,
	}, nil
}
