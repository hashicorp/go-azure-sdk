package mependingaccessreviewinstancestage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePendingAccessReviewInstanceStageClient struct {
	Client *msgraph.Client
}

func NewMePendingAccessReviewInstanceStageClientWithBaseURI(api sdkEnv.Api) (*MePendingAccessReviewInstanceStageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mependingaccessreviewinstancestage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePendingAccessReviewInstanceStageClient: %+v", err)
	}

	return &MePendingAccessReviewInstanceStageClient{
		Client: client,
	}, nil
}
