package transaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransactionClient struct {
	Client *resourcemanager.Client
}

func NewTransactionClientWithBaseURI(sdkApi sdkEnv.Api) (*TransactionClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "transaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TransactionClient: %+v", err)
	}

	return &TransactionClient{
		Client: client,
	}, nil
}
