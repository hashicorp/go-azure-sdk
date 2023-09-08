package userpendingaccessreviewinstancestagedecisioninsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceStageDecisionInsightClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceStageDecisionInsightClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceStageDecisionInsightClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancestagedecisioninsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceStageDecisionInsightClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceStageDecisionInsightClient{
		Client: client,
	}, nil
}
