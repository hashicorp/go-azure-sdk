package serviceprincipallicensedetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalLicenseDetailClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalLicenseDetailClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalLicenseDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipallicensedetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalLicenseDetailClient: %+v", err)
	}

	return &ServicePrincipalLicenseDetailClient{
		Client: client,
	}, nil
}
