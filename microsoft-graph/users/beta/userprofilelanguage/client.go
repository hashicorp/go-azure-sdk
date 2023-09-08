package userprofilelanguage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfileLanguageClient struct {
	Client *msgraph.Client
}

func NewUserProfileLanguageClientWithBaseURI(api sdkEnv.Api) (*UserProfileLanguageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofilelanguage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfileLanguageClient: %+v", err)
	}

	return &UserProfileLanguageClient{
		Client: client,
	}, nil
}
