package mependingaccessreviewinstancedecisioninstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePendingAccessReviewInstanceDecisionInstanceClient struct {
	Client *msgraph.Client
}

func NewMePendingAccessReviewInstanceDecisionInstanceClientWithBaseURI(api sdkEnv.Api) (*MePendingAccessReviewInstanceDecisionInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mependingaccessreviewinstancedecisioninstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePendingAccessReviewInstanceDecisionInstanceClient: %+v", err)
	}

	return &MePendingAccessReviewInstanceDecisionInstanceClient{
		Client: client,
	}, nil
}
