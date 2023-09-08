package setgroupsitetermstoresetparentgroupsetterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreSetParentGroupSetTermClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreSetParentGroupSetTermClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreSetParentGroupSetTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoresetparentgroupsetterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreSetParentGroupSetTermClient: %+v", err)
	}

	return &SetGroupSiteTermStoreSetParentGroupSetTermClient{
		Client: client,
	}, nil
}
