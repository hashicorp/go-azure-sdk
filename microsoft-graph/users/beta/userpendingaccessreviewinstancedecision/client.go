package userpendingaccessreviewinstancedecision

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceDecisionClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceDecisionClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceDecisionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancedecision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceDecisionClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceDecisionClient{
		Client: client,
	}, nil
}
