package serviceprincipaltokenissuancepolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalTokenIssuancePolicyClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalTokenIssuancePolicyClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalTokenIssuancePolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipaltokenissuancepolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalTokenIssuancePolicyClient: %+v", err)
	}

	return &ServicePrincipalTokenIssuancePolicyClient{
		Client: client,
	}, nil
}
