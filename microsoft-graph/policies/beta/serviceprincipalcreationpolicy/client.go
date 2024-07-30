package serviceprincipalcreationpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalCreationPolicyClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalCreationPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*ServicePrincipalCreationPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "serviceprincipalcreationpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalCreationPolicyClient: %+v", err)
	}

	return &ServicePrincipalCreationPolicyClient{
		Client: client,
	}, nil
}
