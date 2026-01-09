package usercredentialusagedetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCredentialUsageDetailClient struct {
	Client *msgraph.Client
}

func NewUserCredentialUsageDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*UserCredentialUsageDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "usercredentialusagedetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCredentialUsageDetailClient: %+v", err)
	}

	return &UserCredentialUsageDetailClient{
		Client: client,
	}, nil
}
