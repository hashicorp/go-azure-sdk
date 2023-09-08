package userpendingaccessreviewinstancedecisioninstancestage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceDecisionInstanceStageClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceDecisionInstanceStageClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceDecisionInstanceStageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancedecisioninstancestage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceDecisionInstanceStageClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceDecisionInstanceStageClient{
		Client: client,
	}, nil
}
