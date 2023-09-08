package userpendingaccessreviewinstancestage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceStageClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceStageClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceStageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancestage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceStageClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceStageClient{
		Client: client,
	}, nil
}
