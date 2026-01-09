package agreementacceptance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgreementAcceptanceClient struct {
	Client *msgraph.Client
}

func NewAgreementAcceptanceClientWithBaseURI(sdkApi sdkEnv.Api) (*AgreementAcceptanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "agreementacceptance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AgreementAcceptanceClient: %+v", err)
	}

	return &AgreementAcceptanceClient{
		Client: client,
	}, nil
}
