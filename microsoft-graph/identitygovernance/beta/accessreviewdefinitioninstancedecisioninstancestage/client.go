package accessreviewdefinitioninstancedecisioninstancestage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceDecisionInstanceStageClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceDecisionInstanceStageClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceDecisionInstanceStageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancedecisioninstancestage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceDecisionInstanceStageClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceDecisionInstanceStageClient{
		Client: client,
	}, nil
}
