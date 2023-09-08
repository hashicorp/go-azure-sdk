package userpendingaccessreviewinstancestagedecisioninstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceStageDecisionInstanceClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceStageDecisionInstanceClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceStageDecisionInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancestagedecisioninstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceStageDecisionInstanceClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceStageDecisionInstanceClient{
		Client: client,
	}, nil
}
