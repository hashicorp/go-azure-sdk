package grouppolicyconfigurationdefinitionvalue

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyConfigurationDefinitionValueClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyConfigurationDefinitionValueClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyConfigurationDefinitionValueClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicyconfigurationdefinitionvalue", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyConfigurationDefinitionValueClient: %+v", err)
	}

	return &GroupPolicyConfigurationDefinitionValueClient{
		Client: client,
	}, nil
}
