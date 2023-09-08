package setgroupsitetermstoregroupsetterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreGroupSetTermClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreGroupSetTermClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreGroupSetTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoregroupsetterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreGroupSetTermClient: %+v", err)
	}

	return &SetGroupSiteTermStoreGroupSetTermClient{
		Client: client,
	}, nil
}
