package groupsitetermstoresetparentgroupsetrelationtoterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetParentGroupSetRelationToTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetParentGroupSetRelationToTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetParentGroupSetRelationToTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetparentgroupsetrelationtoterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetParentGroupSetRelationToTermClient: %+v", err)
	}

	return &GroupSiteTermStoreSetParentGroupSetRelationToTermClient{
		Client: client,
	}, nil
}
