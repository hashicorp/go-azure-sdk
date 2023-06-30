package kafkaconfiguration

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KafkaConfigurationClient struct {
	Client  autorest.Client
	baseUri string
}

func NewKafkaConfigurationClientWithBaseURI(endpoint string) KafkaConfigurationClient {
	return KafkaConfigurationClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
