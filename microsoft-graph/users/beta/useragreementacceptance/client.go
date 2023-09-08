package useragreementacceptance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAgreementAcceptanceClient struct {
	Client *msgraph.Client
}

func NewUserAgreementAcceptanceClientWithBaseURI(api sdkEnv.Api) (*UserAgreementAcceptanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useragreementacceptance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAgreementAcceptanceClient: %+v", err)
	}

	return &UserAgreementAcceptanceClient{
		Client: client,
	}, nil
}
