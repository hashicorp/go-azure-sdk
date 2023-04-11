package workspacesettings

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceSettingsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewWorkspaceSettingsClientWithBaseURI(endpoint string) WorkspaceSettingsClient {
	return WorkspaceSettingsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
