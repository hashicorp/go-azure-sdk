package accessreviewpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewPolicyClient struct {
	Client *msgraph.Client
}

func NewAccessReviewPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*AccessReviewPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "accessreviewpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccessReviewPolicyClient: %+v", err)
	}

	return &AccessReviewPolicyClient{
		Client: client,
	}, nil
}
