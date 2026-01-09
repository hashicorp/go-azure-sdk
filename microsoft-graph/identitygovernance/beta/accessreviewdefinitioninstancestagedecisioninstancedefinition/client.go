package accessreviewdefinitioninstancestagedecisioninstancedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceStageDecisionInstanceDefinitionClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceStageDecisionInstanceDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceStageDecisionInstanceDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancestagedecisioninstancedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceStageDecisionInstanceDefinitionClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceStageDecisionInstanceDefinitionClient{
		Client: client,
	}, nil
}
