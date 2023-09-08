package policydirectoryroleaccessreviewpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyDirectoryRoleAccessReviewPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyDirectoryRoleAccessReviewPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyDirectoryRoleAccessReviewPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policydirectoryroleaccessreviewpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyDirectoryRoleAccessReviewPolicyClient: %+v", err)
	}

	return &PolicyDirectoryRoleAccessReviewPolicyClient{
		Client: client,
	}, nil
}
