package configurationprofileassignments

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationProfileAssignmentsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewConfigurationProfileAssignmentsClientWithBaseURI(endpoint string) ConfigurationProfileAssignmentsClient {
	return ConfigurationProfileAssignmentsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
