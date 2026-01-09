package grouppolicyconfigurationdefinitionvaluepresentationvaluedefinitionvalue

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyConfigurationDefinitionValuePresentationValueDefinitionValueClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyConfigurationDefinitionValuePresentationValueDefinitionValueClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyConfigurationDefinitionValuePresentationValueDefinitionValueClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicyconfigurationdefinitionvaluepresentationvaluedefinitionvalue", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyConfigurationDefinitionValuePresentationValueDefinitionValueClient: %+v", err)
	}

	return &GroupPolicyConfigurationDefinitionValuePresentationValueDefinitionValueClient{
		Client: client,
	}, nil
}
