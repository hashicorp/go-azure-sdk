package managementgroupdiagnosticsettings

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementGroupDiagnosticSettingsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewManagementGroupDiagnosticSettingsClientWithBaseURI(endpoint string) ManagementGroupDiagnosticSettingsClient {
	return ManagementGroupDiagnosticSettingsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
