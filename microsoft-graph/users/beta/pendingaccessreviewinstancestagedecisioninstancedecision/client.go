package pendingaccessreviewinstancestagedecisioninstancedecision

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingAccessReviewInstanceStageDecisionInstanceDecisionClient struct {
	Client *msgraph.Client
}

func NewPendingAccessReviewInstanceStageDecisionInstanceDecisionClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingAccessReviewInstanceStageDecisionInstanceDecisionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingaccessreviewinstancestagedecisioninstancedecision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingAccessReviewInstanceStageDecisionInstanceDecisionClient: %+v", err)
	}

	return &PendingAccessReviewInstanceStageDecisionInstanceDecisionClient{
		Client: client,
	}, nil
}
