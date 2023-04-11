package informationprotectionpolicies

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionPoliciesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewInformationProtectionPoliciesClientWithBaseURI(endpoint string) InformationProtectionPoliciesClient {
	return InformationProtectionPoliciesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
