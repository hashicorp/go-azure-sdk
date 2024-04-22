package validateoperation

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateOperationClient struct {
	Client  autorest.Client
	baseUri string
}

func NewValidateOperationClientWithBaseURI(endpoint string) ValidateOperationClient {
	return ValidateOperationClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
