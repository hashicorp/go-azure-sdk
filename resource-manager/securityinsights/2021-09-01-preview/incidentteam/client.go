package incidentteam

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentTeamClient struct {
	Client  autorest.Client
	baseUri string
}

func NewIncidentTeamClientWithBaseURI(endpoint string) IncidentTeamClient {
	return IncidentTeamClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
