package pendingaccessreviewinstancedecisioninsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingAccessReviewInstanceDecisionInsightClient struct {
	Client *msgraph.Client
}

func NewPendingAccessReviewInstanceDecisionInsightClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingAccessReviewInstanceDecisionInsightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingaccessreviewinstancedecisioninsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingAccessReviewInstanceDecisionInsightClient: %+v", err)
	}

	return &PendingAccessReviewInstanceDecisionInsightClient{
		Client: client,
	}, nil
}
