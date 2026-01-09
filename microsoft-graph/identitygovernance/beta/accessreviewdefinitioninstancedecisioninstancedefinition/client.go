package accessreviewdefinitioninstancedecisioninstancedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceDecisionInstanceDefinitionClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceDecisionInstanceDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceDecisionInstanceDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancedecisioninstancedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceDecisionInstanceDefinitionClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceDecisionInstanceDefinitionClient{
		Client: client,
	}, nil
}
