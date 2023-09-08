package mependingaccessreviewinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePendingAccessReviewInstanceClient struct {
	Client *msgraph.Client
}

func NewMePendingAccessReviewInstanceClientWithBaseURI(api sdkEnv.Api) (*MePendingAccessReviewInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mependingaccessreviewinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePendingAccessReviewInstanceClient: %+v", err)
	}

	return &MePendingAccessReviewInstanceClient{
		Client: client,
	}, nil
}
