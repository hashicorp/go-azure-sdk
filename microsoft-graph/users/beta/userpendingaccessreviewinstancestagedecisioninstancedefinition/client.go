package userpendingaccessreviewinstancestagedecisioninstancedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceStageDecisionInstanceDefinitionClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceStageDecisionInstanceDefinitionClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceStageDecisionInstanceDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancestagedecisioninstancedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceStageDecisionInstanceDefinitionClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceStageDecisionInstanceDefinitionClient{
		Client: client,
	}, nil
}
