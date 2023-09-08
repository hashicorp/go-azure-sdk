package meinformationprotectionbitlockerrecoverykey

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeInformationProtectionBitlockerRecoveryKeyClient struct {
	Client *msgraph.Client
}

func NewMeInformationProtectionBitlockerRecoveryKeyClientWithBaseURI(api sdkEnv.Api) (*MeInformationProtectionBitlockerRecoveryKeyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meinformationprotectionbitlockerrecoverykey", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeInformationProtectionBitlockerRecoveryKeyClient: %+v", err)
	}

	return &MeInformationProtectionBitlockerRecoveryKeyClient{
		Client: client,
	}, nil
}
