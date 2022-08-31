package sourcecontrolsyncjob

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlSyncJobClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSourceControlSyncJobClientWithBaseURI(endpoint string) SourceControlSyncJobClient {
	return SourceControlSyncJobClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
