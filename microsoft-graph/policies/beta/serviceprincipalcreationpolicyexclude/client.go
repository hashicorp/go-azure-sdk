package serviceprincipalcreationpolicyexclude

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalCreationPolicyExcludeClient struct {
	Client *msgraph.Client
}

func NewServicePrincipalCreationPolicyExcludeClientWithBaseURI(sdkApi sdkEnv.Api) (*ServicePrincipalCreationPolicyExcludeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "serviceprincipalcreationpolicyexclude", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServicePrincipalCreationPolicyExcludeClient: %+v", err)
	}

	return &ServicePrincipalCreationPolicyExcludeClient{
		Client: client,
	}, nil
}
