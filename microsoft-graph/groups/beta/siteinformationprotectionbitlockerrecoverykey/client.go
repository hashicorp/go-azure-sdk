package siteinformationprotectionbitlockerrecoverykey

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteInformationProtectionBitlockerRecoveryKeyClient struct {
	Client *msgraph.Client
}

func NewSiteInformationProtectionBitlockerRecoveryKeyClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteInformationProtectionBitlockerRecoveryKeyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteinformationprotectionbitlockerrecoverykey", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteInformationProtectionBitlockerRecoveryKeyClient: %+v", err)
	}

	return &SiteInformationProtectionBitlockerRecoveryKeyClient{
		Client: client,
	}, nil
}
