package useronenotesectiongroupsectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOnenoteSectionGroupSectionPageClient struct {
	Client *msgraph.Client
}

func NewUserOnenoteSectionGroupSectionPageClientWithBaseURI(api sdkEnv.Api) (*UserOnenoteSectionGroupSectionPageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useronenotesectiongroupsectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOnenoteSectionGroupSectionPageClient: %+v", err)
	}

	return &UserOnenoteSectionGroupSectionPageClient{
		Client: client,
	}, nil
}
