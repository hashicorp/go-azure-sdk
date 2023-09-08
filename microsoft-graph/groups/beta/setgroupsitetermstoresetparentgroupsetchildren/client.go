package setgroupsitetermstoresetparentgroupsetchildren

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreSetParentGroupSetChildrenClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreSetParentGroupSetChildrenClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreSetParentGroupSetChildrenClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoresetparentgroupsetchildren", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreSetParentGroupSetChildrenClient: %+v", err)
	}

	return &SetGroupSiteTermStoreSetParentGroupSetChildrenClient{
		Client: client,
	}, nil
}
