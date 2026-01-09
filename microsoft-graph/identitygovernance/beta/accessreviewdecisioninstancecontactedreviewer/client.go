package accessreviewdecisioninstancecontactedreviewer

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDecisionInstanceContactedReviewerClient struct {
	Client *msgraph.Client
}

func NewAccessReviewDecisionInstanceContactedReviewerClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewDecisionInstanceContactedReviewerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewdecisioninstancecontactedreviewer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewDecisionInstanceContactedReviewerClient: %+v", err)
	}

	return &AccessReviewDecisionInstanceContactedReviewerClient{
		Client: client,
	}, nil
}
