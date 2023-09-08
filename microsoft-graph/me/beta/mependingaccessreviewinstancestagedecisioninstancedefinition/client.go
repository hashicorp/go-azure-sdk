package mependingaccessreviewinstancestagedecisioninstancedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePendingAccessReviewInstanceStageDecisionInstanceDefinitionClient struct {
	Client *msgraph.Client
}

func NewMePendingAccessReviewInstanceStageDecisionInstanceDefinitionClientWithBaseURI(api sdkEnv.Api) (*MePendingAccessReviewInstanceStageDecisionInstanceDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mependingaccessreviewinstancestagedecisioninstancedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePendingAccessReviewInstanceStageDecisionInstanceDefinitionClient: %+v", err)
	}

	return &MePendingAccessReviewInstanceStageDecisionInstanceDefinitionClient{
		Client: client,
	}, nil
}
