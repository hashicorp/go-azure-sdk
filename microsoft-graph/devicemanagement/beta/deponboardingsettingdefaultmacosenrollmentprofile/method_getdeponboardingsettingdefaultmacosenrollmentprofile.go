package deponboardingsettingdefaultmacosenrollmentprofile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDepOnboardingSettingDefaultMacOsEnrollmentProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DepMacOSEnrollmentProfile
}

type GetDepOnboardingSettingDefaultMacOsEnrollmentProfileOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetDepOnboardingSettingDefaultMacOsEnrollmentProfileOperationOptions() GetDepOnboardingSettingDefaultMacOsEnrollmentProfileOperationOptions {
	return GetDepOnboardingSettingDefaultMacOsEnrollmentProfileOperationOptions{}
}

func (o GetDepOnboardingSettingDefaultMacOsEnrollmentProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDepOnboardingSettingDefaultMacOsEnrollmentProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDepOnboardingSettingDefaultMacOsEnrollmentProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDepOnboardingSettingDefaultMacOsEnrollmentProfile - Get defaultMacOsEnrollmentProfile from deviceManagement.
// Default MacOs Enrollment Profile
func (c DepOnboardingSettingDefaultMacOsEnrollmentProfileClient) GetDepOnboardingSettingDefaultMacOsEnrollmentProfile(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingId, options GetDepOnboardingSettingDefaultMacOsEnrollmentProfileOperationOptions) (result GetDepOnboardingSettingDefaultMacOsEnrollmentProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/defaultMacOsEnrollmentProfile", id.ID()),
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

	var model beta.DepMacOSEnrollmentProfile
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
