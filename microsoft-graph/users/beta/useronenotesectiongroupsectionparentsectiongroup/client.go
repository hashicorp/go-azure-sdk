package useronenotesectiongroupsectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteSectionGroupSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteSectionGroupSectionParentSectionGroupClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteSectionGroupSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotesectiongroupsectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteSectionGroupSectionParentSectionGroupClient: %+v", err)
	}

	return &UserOnenoteSectionGroupSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
