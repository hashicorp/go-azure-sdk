package paymentmethods

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PaymentMethodsClient struct {
	Client *resourcemanager.Client
}

func NewPaymentMethodsClientWithBaseURI(sdkApi sdkEnv.Api) (*PaymentMethodsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "paymentmethods", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PaymentMethodsClient: %+v", err)
	}

	return &PaymentMethodsClient{
		Client: client,
	}, nil
}
