package userinformationprotectionpolicylabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInformationProtectionPolicyLabelClient struct {
	Client *msgraph.Client
}

func NewUserInformationProtectionPolicyLabelClientWithBaseURI(api sdkEnv.Api) (*UserInformationProtectionPolicyLabelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinformationprotectionpolicylabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInformationProtectionPolicyLabelClient: %+v", err)
	}

	return &UserInformationProtectionPolicyLabelClient{
		Client: client,
	}, nil
}
