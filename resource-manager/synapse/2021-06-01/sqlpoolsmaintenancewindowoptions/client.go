package sqlpoolsmaintenancewindowoptions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsMaintenanceWindowOptionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSqlPoolsMaintenanceWindowOptionsClientWithBaseURI(endpoint string) SqlPoolsMaintenanceWindowOptionsClient {
	return SqlPoolsMaintenanceWindowOptionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
