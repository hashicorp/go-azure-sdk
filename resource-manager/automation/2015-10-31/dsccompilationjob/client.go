package dsccompilationjob

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscCompilationJobClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDscCompilationJobClientWithBaseURI(endpoint string) DscCompilationJobClient {
	return DscCompilationJobClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
