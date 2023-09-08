package userprofileanniversary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfileAnniversaryClient struct {
	Client *msgraph.Client
}

func NewUserProfileAnniversaryClientWithBaseURI(api sdkEnv.Api) (*UserProfileAnniversaryClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofileanniversary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfileAnniversaryClient: %+v", err)
	}

	return &UserProfileAnniversaryClient{
		Client: client,
	}, nil
}
