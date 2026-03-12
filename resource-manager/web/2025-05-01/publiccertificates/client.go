package publiccertificates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublicCertificatesClient struct {
	Client *resourcemanager.Client
}

func NewPublicCertificatesClientWithBaseURI(sdkApi sdkEnv.Api) (*PublicCertificatesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "publiccertificates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PublicCertificatesClient: %+v", err)
	}

	return &PublicCertificatesClient{
		Client: client,
	}, nil
}
