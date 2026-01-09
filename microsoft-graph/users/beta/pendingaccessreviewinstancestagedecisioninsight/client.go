package pendingaccessreviewinstancestagedecisioninsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingAccessReviewInstanceStageDecisionInsightClient struct {
	Client *msgraph.Client
}

func NewPendingAccessReviewInstanceStageDecisionInsightClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingAccessReviewInstanceStageDecisionInsightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingaccessreviewinstancestagedecisioninsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingAccessReviewInstanceStageDecisionInsightClient: %+v", err)
	}

	return &PendingAccessReviewInstanceStageDecisionInsightClient{
		Client: client,
	}, nil
}
