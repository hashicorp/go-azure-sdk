package sqlpoolssqlpooluseractivities

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsSqlPoolUserActivitiesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSqlPoolsSqlPoolUserActivitiesClientWithBaseURI(endpoint string) SqlPoolsSqlPoolUserActivitiesClient {
	return SqlPoolsSqlPoolUserActivitiesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
