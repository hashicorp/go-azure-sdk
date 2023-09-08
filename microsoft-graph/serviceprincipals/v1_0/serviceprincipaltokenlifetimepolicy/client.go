package serviceprincipaltokenlifetimepolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalTokenLifetimePolicyClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalTokenLifetimePolicyClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalTokenLifetimePolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipaltokenlifetimepolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalTokenLifetimePolicyClient: %+v", err)
	}

	return &ServicePrincipalTokenLifetimePolicyClient{
		Client: client,
	}, nil
}
