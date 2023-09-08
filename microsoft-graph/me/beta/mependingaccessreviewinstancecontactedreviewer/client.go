package mependingaccessreviewinstancecontactedreviewer

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePendingAccessReviewInstanceContactedReviewerClient struct {
	Client *msgraph.Client
}

func NewMePendingAccessReviewInstanceContactedReviewerClientWithBaseURI(api sdkEnv.Api) (*MePendingAccessReviewInstanceContactedReviewerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mependingaccessreviewinstancecontactedreviewer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePendingAccessReviewInstanceContactedReviewerClient: %+v", err)
	}

	return &MePendingAccessReviewInstanceContactedReviewerClient{
		Client: client,
	}, nil
}
