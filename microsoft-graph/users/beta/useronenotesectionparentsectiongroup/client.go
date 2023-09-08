package useronenotesectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteSectionParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotesectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteSectionParentSectionGroupClient: %+v", err)
	}

	return &UserOnenoteSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
