package incidentcomments

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentCommentsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewIncidentCommentsClientWithBaseURI(endpoint string) IncidentCommentsClient {
	return IncidentCommentsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
