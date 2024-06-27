package billingrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingRequestClient struct {
	Client *resourcemanager.Client
}

func NewBillingRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingRequestClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "billingrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingRequestClient: %+v", err)
	}

	return &BillingRequestClient{
		Client: client,
	}, nil
}
