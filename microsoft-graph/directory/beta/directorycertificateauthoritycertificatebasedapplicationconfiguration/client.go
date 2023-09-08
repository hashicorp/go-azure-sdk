package directorycertificateauthoritycertificatebasedapplicationconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationClient struct {
	Client *msgraph.Client
}

func NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationClientWithBaseURI(api sdkEnv.Api) (*DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directorycertificateauthoritycertificatebasedapplicationconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationClient: %+v", err)
	}

	return &DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationClient{
		Client: client,
	}, nil
}
