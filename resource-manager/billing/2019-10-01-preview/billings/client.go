package billings

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingsClient struct {
	Client *resourcemanager.Client
}

func NewBillingsClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "billings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingsClient: %+v", err)
	}

	return &BillingsClient{
		Client: client,
	}, nil
}
