package informationprotectionbitlockerrecoverykey

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionBitlockerRecoveryKeyClient struct {
	Client *msgraph.Client
}

func NewInformationProtectionBitlockerRecoveryKeyClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionBitlockerRecoveryKeyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "informationprotectionbitlockerrecoverykey", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionBitlockerRecoveryKeyClient: %+v", err)
	}

	return &InformationProtectionBitlockerRecoveryKeyClient{
		Client: client,
	}, nil
}
