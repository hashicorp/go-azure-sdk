package pendingaccessreviewinstancecontactedreviewer

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingAccessReviewInstanceContactedReviewerClient struct {
	Client *msgraph.Client
}

func NewPendingAccessReviewInstanceContactedReviewerClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingAccessReviewInstanceContactedReviewerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingaccessreviewinstancecontactedreviewer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingAccessReviewInstanceContactedReviewerClient: %+v", err)
	}

	return &PendingAccessReviewInstanceContactedReviewerClient{
		Client: client,
	}, nil
}
