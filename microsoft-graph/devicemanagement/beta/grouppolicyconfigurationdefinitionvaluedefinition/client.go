package grouppolicyconfigurationdefinitionvaluedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyConfigurationDefinitionValueDefinitionClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyConfigurationDefinitionValueDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyConfigurationDefinitionValueDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicyconfigurationdefinitionvaluedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyConfigurationDefinitionValueDefinitionClient: %+v", err)
	}

	return &GroupPolicyConfigurationDefinitionValueDefinitionClient{
		Client: client,
	}, nil
}
