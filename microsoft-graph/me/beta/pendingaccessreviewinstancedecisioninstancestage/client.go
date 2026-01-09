package pendingaccessreviewinstancedecisioninstancestage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingAccessReviewInstanceDecisionInstanceStageClient struct {
	Client *msgraph.Client
}

func NewPendingAccessReviewInstanceDecisionInstanceStageClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingAccessReviewInstanceDecisionInstanceStageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingaccessreviewinstancedecisioninstancestage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingAccessReviewInstanceDecisionInstanceStageClient: %+v", err)
	}

	return &PendingAccessReviewInstanceDecisionInstanceStageClient{
		Client: client,
	}, nil
}
