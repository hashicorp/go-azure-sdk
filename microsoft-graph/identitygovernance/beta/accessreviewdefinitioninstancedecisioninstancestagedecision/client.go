package accessreviewdefinitioninstancedecisioninstancestagedecision

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceDecisionInstanceStageDecisionClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceDecisionInstanceStageDecisionClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceDecisionInstanceStageDecisionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancedecisioninstancestagedecision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceDecisionInstanceStageDecisionClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceDecisionInstanceStageDecisionClient{
		Client: client,
	}, nil
}
