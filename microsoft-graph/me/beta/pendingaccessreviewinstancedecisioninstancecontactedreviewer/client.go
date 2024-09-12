package pendingaccessreviewinstancedecisioninstancecontactedreviewer

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingAccessReviewInstanceDecisionInstanceContactedReviewerClient struct {
	Client *msgraph.Client
}

func NewPendingAccessReviewInstanceDecisionInstanceContactedReviewerClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingAccessReviewInstanceDecisionInstanceContactedReviewerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingaccessreviewinstancedecisioninstancecontactedreviewer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingAccessReviewInstanceDecisionInstanceContactedReviewerClient: %+v", err)
	}

	return &PendingAccessReviewInstanceDecisionInstanceContactedReviewerClient{
		Client: client,
	}, nil
}
