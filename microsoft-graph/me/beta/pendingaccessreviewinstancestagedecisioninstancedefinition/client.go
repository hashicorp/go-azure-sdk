package pendingaccessreviewinstancestagedecisioninstancedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingAccessReviewInstanceStageDecisionInstanceDefinitionClient struct {
	Client *msgraph.Client
}

func NewPendingAccessReviewInstanceStageDecisionInstanceDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingAccessReviewInstanceStageDecisionInstanceDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingaccessreviewinstancestagedecisioninstancedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingAccessReviewInstanceStageDecisionInstanceDefinitionClient: %+v", err)
	}

	return &PendingAccessReviewInstanceStageDecisionInstanceDefinitionClient{
		Client: client,
	}, nil
}
