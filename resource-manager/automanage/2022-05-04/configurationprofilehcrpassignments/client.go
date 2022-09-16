package configurationprofilehcrpassignments

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationProfileHCRPAssignmentsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewConfigurationProfileHCRPAssignmentsClientWithBaseURI(endpoint string) ConfigurationProfileHCRPAssignmentsClient {
	return ConfigurationProfileHCRPAssignmentsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
