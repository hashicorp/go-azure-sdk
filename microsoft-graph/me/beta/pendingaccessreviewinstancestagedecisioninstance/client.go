package pendingaccessreviewinstancestagedecisioninstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingAccessReviewInstanceStageDecisionInstanceClient struct {
	Client *msgraph.Client
}

func NewPendingAccessReviewInstanceStageDecisionInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingAccessReviewInstanceStageDecisionInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingaccessreviewinstancestagedecisioninstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingAccessReviewInstanceStageDecisionInstanceClient: %+v", err)
	}

	return &PendingAccessReviewInstanceStageDecisionInstanceClient{
		Client: client,
	}, nil
}
