package accessreviewdefinitioninstancestagedecisioninstancecontactedreviewer

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdefinitioninstancestagedecisioninstancecontactedreviewer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerClient: %+v", err)
	}

	return &AccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerClient{
		Client: client,
	}, nil
}
