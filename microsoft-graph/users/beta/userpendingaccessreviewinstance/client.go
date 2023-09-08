package userpendingaccessreviewinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPendingAccessReviewInstanceClient struct {
	Client *msgraph.Client
}

func NewUserPendingAccessReviewInstanceClientWithBaseURI(api sdkEnv.Api) (*UserPendingAccessReviewInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userpendingaccessreviewinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPendingAccessReviewInstanceClient: %+v", err)
	}

	return &UserPendingAccessReviewInstanceClient{
		Client: client,
	}, nil
}
