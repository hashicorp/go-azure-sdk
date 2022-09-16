package sqlpoolsmaintenancewindows

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsMaintenanceWindowsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSqlPoolsMaintenanceWindowsClientWithBaseURI(endpoint string) SqlPoolsMaintenanceWindowsClient {
	return SqlPoolsMaintenanceWindowsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
