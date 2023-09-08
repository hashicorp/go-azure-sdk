package setgroupsitetermstoresetparentgroupsetrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreSetParentGroupSetRelationClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreSetParentGroupSetRelationClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreSetParentGroupSetRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoresetparentgroupsetrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreSetParentGroupSetRelationClient: %+v", err)
	}

	return &SetGroupSiteTermStoreSetParentGroupSetRelationClient{
		Client: client,
	}, nil
}
