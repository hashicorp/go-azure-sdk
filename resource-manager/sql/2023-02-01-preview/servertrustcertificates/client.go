package servertrustcertificates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerTrustCertificatesClient struct {
	Client *resourcemanager.Client
}

func NewServerTrustCertificatesClientWithBaseURI(sdkApi sdkEnv.Api) (*ServerTrustCertificatesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "servertrustcertificates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerTrustCertificatesClient: %+v", err)
	}

	return &ServerTrustCertificatesClient{
		Client: client,
	}, nil
}
