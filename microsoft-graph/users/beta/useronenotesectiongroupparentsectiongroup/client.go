package useronenotesectiongroupparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteSectionGroupParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteSectionGroupParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteSectionGroupParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotesectiongroupparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteSectionGroupParentSectionGroupClient: %+v", err)
	}

	return &UserOnenoteSectionGroupParentSectionGroupClient{
		Client: client,
	}, nil
}
