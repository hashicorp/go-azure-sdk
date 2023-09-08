package userpendingaccessreviewinstancestagedecisioninstancedecisioninsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancestagedecisioninstancedecisioninsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceStageDecisionInstanceDecisionInsightClient{
		Client: client,
	}, nil
}
