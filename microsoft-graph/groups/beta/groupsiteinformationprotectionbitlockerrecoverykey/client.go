package groupsiteinformationprotectionbitlockerrecoverykey

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteInformationProtectionBitlockerRecoveryKeyClient struct {
	Client *msgraph.Client
}

func NewGroupSiteInformationProtectionBitlockerRecoveryKeyClientWithBaseURI(api sdkEnv.Api) (*GroupSiteInformationProtectionBitlockerRecoveryKeyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteinformationprotectionbitlockerrecoverykey", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteInformationProtectionBitlockerRecoveryKeyClient: %+v", err)
	}

	return &GroupSiteInformationProtectionBitlockerRecoveryKeyClient{
		Client: client,
	}, nil
}
