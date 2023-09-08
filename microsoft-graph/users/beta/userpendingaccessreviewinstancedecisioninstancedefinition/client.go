package userpendingaccessreviewinstancedecisioninstancedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceDecisionInstanceDefinitionClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceDecisionInstanceDefinitionClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceDecisionInstanceDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancedecisioninstancedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceDecisionInstanceDefinitionClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceDecisionInstanceDefinitionClient{
		Client: client,
	}, nil
}
