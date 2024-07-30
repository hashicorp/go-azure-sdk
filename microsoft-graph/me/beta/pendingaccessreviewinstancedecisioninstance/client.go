package pendingaccessreviewinstancedecisioninstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingAccessReviewInstanceDecisionInstanceClient struct {
	Client *msgraph.Client
}

func NewPendingAccessReviewInstanceDecisionInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingAccessReviewInstanceDecisionInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingaccessreviewinstancedecisioninstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingAccessReviewInstanceDecisionInstanceClient: %+v", err)
	}

	return &PendingAccessReviewInstanceDecisionInstanceClient{
		Client: client,
	}, nil
}
