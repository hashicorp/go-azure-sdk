package grouppolicyconfigurationdefinitionvaluepresentationvaluepresentation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyConfigurationDefinitionValuePresentationValuePresentationClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyConfigurationDefinitionValuePresentationValuePresentationClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyConfigurationDefinitionValuePresentationValuePresentationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicyconfigurationdefinitionvaluepresentationvaluepresentation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyConfigurationDefinitionValuePresentationValuePresentationClient: %+v", err)
	}

	return &GroupPolicyConfigurationDefinitionValuePresentationValuePresentationClient{
		Client: client,
	}, nil
}
