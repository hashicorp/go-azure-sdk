package v2022_10_31_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/data-plane/iotcentral/2022-10-31-preview/iotcentrals"
	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

type Client struct {
	Iotcentrals *iotcentrals.IotcentralsClient
}

func NewClient(configureFunc func(c *dataplane.Client)) (*Client, error) {
	iotcentralsClient, err := iotcentrals.NewIotcentralsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Iotcentrals client: %+v", err)
	}
	configureFunc(iotcentralsClient.Client)

	return &Client{
		Iotcentrals: iotcentralsClient,
	}, nil
}
