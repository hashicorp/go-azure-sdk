package pendingaccessreviewinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingAccessReviewInstanceClient struct {
	Client *msgraph.Client
}

func NewPendingAccessReviewInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingAccessReviewInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingaccessreviewinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingAccessReviewInstanceClient: %+v", err)
	}

	return &PendingAccessReviewInstanceClient{
		Client: client,
	}, nil
}
