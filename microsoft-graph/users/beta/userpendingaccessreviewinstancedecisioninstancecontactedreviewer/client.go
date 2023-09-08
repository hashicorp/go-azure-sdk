package userpendingaccessreviewinstancedecisioninstancecontactedreviewer

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceDecisionInstanceContactedReviewerClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancedecisioninstancecontactedreviewer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceDecisionInstanceContactedReviewerClient{
		Client: client,
	}, nil
}
