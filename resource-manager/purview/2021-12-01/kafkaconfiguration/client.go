package kafkaconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KafkaConfigurationClient struct {
	Client *resourcemanager.Client
}

func NewKafkaConfigurationClientWithBaseURI(api environments.Api) (*KafkaConfigurationClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "kafkaconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating KafkaConfigurationClient: %+v", err)
	}

	return &KafkaConfigurationClient{
		Client: client,
	}, nil
}
