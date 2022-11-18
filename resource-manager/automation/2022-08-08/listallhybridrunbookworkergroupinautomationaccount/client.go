package listallhybridrunbookworkergroupinautomationaccount

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAllHybridRunbookWorkerGroupInAutomationAccountClient struct {
	Client  autorest.Client
	baseUri string
}

func NewListAllHybridRunbookWorkerGroupInAutomationAccountClientWithBaseURI(endpoint string) ListAllHybridRunbookWorkerGroupInAutomationAccountClient {
	return ListAllHybridRunbookWorkerGroupInAutomationAccountClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
