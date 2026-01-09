package deponboardingsettingenrollmentprofile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetDepOnboardingSettingEnrollmentProfileDefaultProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetDepOnboardingSettingEnrollmentProfileDefaultProfileOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSetDepOnboardingSettingEnrollmentProfileDefaultProfileOperationOptions() SetDepOnboardingSettingEnrollmentProfileDefaultProfileOperationOptions {
	return SetDepOnboardingSettingEnrollmentProfileDefaultProfileOperationOptions{}
}

func (o SetDepOnboardingSettingEnrollmentProfileDefaultProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetDepOnboardingSettingEnrollmentProfileDefaultProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetDepOnboardingSettingEnrollmentProfileDefaultProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetDepOnboardingSettingEnrollmentProfileDefaultProfile - Invoke action setDefaultProfile
func (c DepOnboardingSettingEnrollmentProfileClient) SetDepOnboardingSettingEnrollmentProfileDefaultProfile(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingIdEnrollmentProfileId, options SetDepOnboardingSettingEnrollmentProfileDefaultProfileOperationOptions) (result SetDepOnboardingSettingEnrollmentProfileDefaultProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/setDefaultProfile", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
