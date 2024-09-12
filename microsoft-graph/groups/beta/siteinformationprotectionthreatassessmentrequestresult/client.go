package siteinformationprotectionthreatassessmentrequestresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteInformationProtectionThreatAssessmentRequestResultClient struct {
	Client *msgraph.Client
}

func NewSiteInformationProtectionThreatAssessmentRequestResultClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteInformationProtectionThreatAssessmentRequestResultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteinformationprotectionthreatassessmentrequestresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteInformationProtectionThreatAssessmentRequestResultClient: %+v", err)
	}

	return &SiteInformationProtectionThreatAssessmentRequestResultClient{
		Client: client,
	}, nil
}
