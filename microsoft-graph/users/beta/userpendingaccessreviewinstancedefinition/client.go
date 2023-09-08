package userpendingaccessreviewinstancedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceDefinitionClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceDefinitionClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstancedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceDefinitionClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceDefinitionClient{
		Client: client,
	}, nil
}
