package mependingaccessreviewinstancedecisioninstancestagedecisioninsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient struct {
	Client *msgraph.Client
}

func NewMePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClientWithBaseURI(api sdkEnv.Api) (*MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mependingaccessreviewinstancedecisioninstancestagedecisioninsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient: %+v", err)
	}

	return &MePendingAccessReviewInstanceDecisionInstanceStageDecisionInsightClient{
		Client: client,
	}, nil
}
