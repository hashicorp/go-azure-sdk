package userpendingaccessreviewinstancedecisioninsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceDecisionInsightClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceDecisionInsightClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceDecisionInsightClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancedecisioninsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceDecisionInsightClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceDecisionInsightClient{
		Client: client,
	}, nil
}
