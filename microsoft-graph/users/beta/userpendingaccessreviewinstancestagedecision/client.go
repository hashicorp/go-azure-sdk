package userpendingaccessreviewinstancestagedecision

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceStageDecisionClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceStageDecisionClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceStageDecisionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancestagedecision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceStageDecisionClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceStageDecisionClient{
		Client: client,
	}, nil
}
