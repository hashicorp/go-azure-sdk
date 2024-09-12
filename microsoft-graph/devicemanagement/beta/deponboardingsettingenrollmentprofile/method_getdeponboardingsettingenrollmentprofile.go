package deponboardingsettingenrollmentprofile

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDepOnboardingSettingEnrollmentProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.EnrollmentProfile
}

type GetDepOnboardingSettingEnrollmentProfileOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDepOnboardingSettingEnrollmentProfileOperationOptions() GetDepOnboardingSettingEnrollmentProfileOperationOptions {
	return GetDepOnboardingSettingEnrollmentProfileOperationOptions{}
}

func (o GetDepOnboardingSettingEnrollmentProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDepOnboardingSettingEnrollmentProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDepOnboardingSettingEnrollmentProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDepOnboardingSettingEnrollmentProfile - Get enrollmentProfiles from deviceManagement. The enrollment profiles.
func (c DepOnboardingSettingEnrollmentProfileClient) GetDepOnboardingSettingEnrollmentProfile(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingIdEnrollmentProfileId, options GetDepOnboardingSettingEnrollmentProfileOperationOptions) (result GetDepOnboardingSettingEnrollmentProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalEnrollmentProfileImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
