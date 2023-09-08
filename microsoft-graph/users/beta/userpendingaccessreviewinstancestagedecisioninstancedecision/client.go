package userpendingaccessreviewinstancestagedecisioninstancedecision

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceStageDecisionInstanceDecisionClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancestagedecisioninstancedecision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceStageDecisionInstanceDecisionClient{
		Client: client,
	}, nil
}
