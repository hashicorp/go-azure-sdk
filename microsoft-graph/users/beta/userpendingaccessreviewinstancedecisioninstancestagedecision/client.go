package userpendingaccessreviewinstancedecisioninstancestagedecision

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceDecisionInstanceStageDecisionClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancedecisioninstancestagedecision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceDecisionInstanceStageDecisionClient{
		Client: client,
	}, nil
}
