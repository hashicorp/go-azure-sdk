package userinformationprotectionbitlockerrecoverykey

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInformationProtectionBitlockerRecoveryKeyClient struct {
	Client *msgraph.Client
}

func NewUserInformationProtectionBitlockerRecoveryKeyClientWithBaseURI(api sdkEnv.Api) (*UserInformationProtectionBitlockerRecoveryKeyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinformationprotectionbitlockerrecoverykey", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInformationProtectionBitlockerRecoveryKeyClient: %+v", err)
	}

	return &UserInformationProtectionBitlockerRecoveryKeyClient{
		Client: client,
	}, nil
}
