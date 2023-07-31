package namespacesauthorizationrule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespacesAuthorizationRuleClient struct {
	Client *resourcemanager.Client
}

func NewNamespacesAuthorizationRuleClientWithBaseURI(api sdkEnv.Api) (*NamespacesAuthorizationRuleClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "namespacesauthorizationrule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NamespacesAuthorizationRuleClient: %+v", err)
	}

	return &NamespacesAuthorizationRuleClient{
		Client: client,
	}, nil
}
