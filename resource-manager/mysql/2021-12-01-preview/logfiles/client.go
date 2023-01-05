package logfiles

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogFilesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewLogFilesClientWithBaseURI(endpoint string) LogFilesClient {
	return LogFilesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
