package pendingaccessreviewinstancestage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingAccessReviewInstanceStageClient struct {
	Client *msgraph.Client
}

func NewPendingAccessReviewInstanceStageClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingAccessReviewInstanceStageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingaccessreviewinstancestage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingAccessReviewInstanceStageClient: %+v", err)
	}

	return &PendingAccessReviewInstanceStageClient{
		Client: client,
	}, nil
}
