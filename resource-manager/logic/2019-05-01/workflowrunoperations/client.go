package workflowrunoperations

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowRunOperationsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewWorkflowRunOperationsClientWithBaseURI(endpoint string) WorkflowRunOperationsClient {
	return WorkflowRunOperationsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
