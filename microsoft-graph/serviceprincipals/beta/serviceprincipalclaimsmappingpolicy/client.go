package serviceprincipalclaimsmappingpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalClaimsMappingPolicyClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalClaimsMappingPolicyClientWithBaseURI(api sdkEnv.Api) (*ServicePrincipalClaimsMappingPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "serviceprincipalclaimsmappingpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalClaimsMappingPolicyClient: %+v", err)
	}

	return &ServicePrincipalClaimsMappingPolicyClient{
		Client: client,
	}, nil
}
