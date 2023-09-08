package mependingaccessreviewinstancestagedecisioninsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePendingAccessReviewInstanceStageDecisionInsightClient struct {
	Client *msgraph.Client
}

func NewMePendingAccessReviewInstanceStageDecisionInsightClientWithBaseURI(api sdkEnv.Api) (*MePendingAccessReviewInstanceStageDecisionInsightClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mependingaccessreviewinstancestagedecisioninsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePendingAccessReviewInstanceStageDecisionInsightClient: %+v", err)
	}

	return &MePendingAccessReviewInstanceStageDecisionInsightClient{
		Client: client,
	}, nil
}
