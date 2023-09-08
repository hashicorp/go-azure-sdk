package mependingaccessreviewinstancedecisioninsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePendingAccessReviewInstanceDecisionInsightClient struct {
	Client *msgraph.Client
}

func NewMePendingAccessReviewInstanceDecisionInsightClientWithBaseURI(api sdkEnv.Api) (*MePendingAccessReviewInstanceDecisionInsightClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mependingaccessreviewinstancedecisioninsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePendingAccessReviewInstanceDecisionInsightClient: %+v", err)
	}

	return &MePendingAccessReviewInstanceDecisionInsightClient{
		Client: client,
	}, nil
}
