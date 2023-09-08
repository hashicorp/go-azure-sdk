package policyaccessreviewpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAccessReviewPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyAccessReviewPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyAccessReviewPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyaccessreviewpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyAccessReviewPolicyClient: %+v", err)
	}

	return &PolicyAccessReviewPolicyClient{
		Client: client,
	}, nil
}
