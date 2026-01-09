package pendingexternaluserprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PendingExternalUserProfileClient struct {
	Client *msgraph.Client
}

func NewPendingExternalUserProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*PendingExternalUserProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "pendingexternaluserprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PendingExternalUserProfileClient: %+v", err)
	}

	return &PendingExternalUserProfileClient{
		Client: client,
	}, nil
}
