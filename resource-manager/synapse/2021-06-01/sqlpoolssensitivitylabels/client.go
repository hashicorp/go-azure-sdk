package sqlpoolssensitivitylabels

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsSensitivityLabelsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSqlPoolsSensitivityLabelsClientWithBaseURI(endpoint string) SqlPoolsSensitivityLabelsClient {
	return SqlPoolsSensitivityLabelsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
