package sensitivitylabels

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabelsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSensitivityLabelsClientWithBaseURI(endpoint string) SensitivityLabelsClient {
	return SensitivityLabelsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
