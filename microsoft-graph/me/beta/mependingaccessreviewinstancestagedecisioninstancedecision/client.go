package mependingaccessreviewinstancestagedecisioninstancedecision

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient struct {
	Client *msgraph.Client
}

func NewMePendingAccessReviewInstanceStageDecisionInstanceDecisionClientWithBaseURI(api sdkEnv.Api) (*MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mependingaccessreviewinstancestagedecisioninstancedecision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient: %+v", err)
	}

	return &MePendingAccessReviewInstanceStageDecisionInstanceDecisionClient{
		Client: client,
	}, nil
}
