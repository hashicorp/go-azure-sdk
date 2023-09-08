package groupsitetermstoregroupsetparentgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetParentGroupClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetParentGroupClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetParentGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsetparentgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetParentGroupClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetParentGroupClient{
		Client: client,
	}, nil
}
