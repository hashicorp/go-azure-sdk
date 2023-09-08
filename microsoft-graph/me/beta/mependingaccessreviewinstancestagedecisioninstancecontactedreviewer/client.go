package mependingaccessreviewinstancestagedecisioninstancecontactedreviewer

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient struct {
	Client *msgraph.Client
}

func NewMePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClientWithBaseURI(api sdkEnv.Api) (*MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mependingaccessreviewinstancestagedecisioninstancecontactedreviewer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient: %+v", err)
	}

	return &MePendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient{
		Client: client,
	}, nil
}
