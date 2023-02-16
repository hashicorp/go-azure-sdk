package logger

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoggerClient struct {
	Client  autorest.Client
	baseUri string
}

func NewLoggerClientWithBaseURI(endpoint string) LoggerClient {
	return LoggerClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
