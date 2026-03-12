package publiccertificateoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublicCertificateOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewPublicCertificateOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PublicCertificateOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "publiccertificateoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PublicCertificateOperationGroupClient: %+v", err)
	}

	return &PublicCertificateOperationGroupClient{
		Client: client,
	}, nil
}
