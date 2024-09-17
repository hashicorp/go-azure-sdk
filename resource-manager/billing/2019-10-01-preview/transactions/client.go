package transactions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransactionsClient struct {
	Client *resourcemanager.Client
}

func NewTransactionsClientWithBaseURI(sdkApi sdkEnv.Api) (*TransactionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "transactions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TransactionsClient: %+v", err)
	}

	return &TransactionsClient{
		Client: client,
	}, nil
}
