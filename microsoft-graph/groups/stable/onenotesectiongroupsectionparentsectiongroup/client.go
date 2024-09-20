package onenotesectiongroupsectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionGroupSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewOnenoteSectionGroupSectionParentSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*OnenoteSectionGroupSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "onenotesectiongroupsectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OnenoteSectionGroupSectionParentSectionGroupClient: %+v", err)
	}

	return &OnenoteSectionGroupSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
