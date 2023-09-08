package groupsitetermstoregroupsetchildrenchildren

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetChildrenChildrenClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetChildrenChildrenClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetChildrenChildrenClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsetchildrenchildren", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetChildrenChildrenClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetChildrenChildrenClient{
		Client: client,
	}, nil
}
