package armtemplates

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArmTemplatesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewArmTemplatesClientWithBaseURI(endpoint string) ArmTemplatesClient {
	return ArmTemplatesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
