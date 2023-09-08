package userprofilecertification

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfileCertificationClient struct {
	Client *msgraph.Client
}

func NewUserProfileCertificationClientWithBaseURI(api sdkEnv.Api) (*UserProfileCertificationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofilecertification", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfileCertificationClient: %+v", err)
	}

	return &UserProfileCertificationClient{
		Client: client,
	}, nil
}
