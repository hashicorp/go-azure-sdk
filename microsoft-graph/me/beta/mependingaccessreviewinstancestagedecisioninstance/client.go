package mependingaccessreviewinstancestagedecisioninstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePendingAccessReviewInstanceStageDecisionInstanceClient struct {
	Client *msgraph.Client
}

func NewMePendingAccessReviewInstanceStageDecisionInstanceClientWithBaseURI(api sdkEnv.Api) (*MePendingAccessReviewInstanceStageDecisionInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mependingaccessreviewinstancestagedecisioninstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePendingAccessReviewInstanceStageDecisionInstanceClient: %+v", err)
	}

	return &MePendingAccessReviewInstanceStageDecisionInstanceClient{
		Client: client,
	}, nil
}
