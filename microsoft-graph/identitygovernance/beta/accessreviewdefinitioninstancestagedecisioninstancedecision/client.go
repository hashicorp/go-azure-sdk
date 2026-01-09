package accessreviewdefinitioninstancestagedecisioninstancedecision

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceStageDecisionInstanceDecisionClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceStageDecisionInstanceDecisionClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceStageDecisionInstanceDecisionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancestagedecisioninstancedecision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceStageDecisionInstanceDecisionClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceStageDecisionInstanceDecisionClient{
		Client: client,
	}, nil
}
