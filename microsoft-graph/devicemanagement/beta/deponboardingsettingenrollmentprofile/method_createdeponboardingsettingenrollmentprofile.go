package deponboardingsettingenrollmentprofile

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDepOnboardingSettingEnrollmentProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.EnrollmentProfile
}

type CreateDepOnboardingSettingEnrollmentProfileOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDepOnboardingSettingEnrollmentProfileOperationOptions() CreateDepOnboardingSettingEnrollmentProfileOperationOptions {
	return CreateDepOnboardingSettingEnrollmentProfileOperationOptions{}
}

func (o CreateDepOnboardingSettingEnrollmentProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDepOnboardingSettingEnrollmentProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDepOnboardingSettingEnrollmentProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDepOnboardingSettingEnrollmentProfile - Create new navigation property to enrollmentProfiles for
// deviceManagement
func (c DepOnboardingSettingEnrollmentProfileClient) CreateDepOnboardingSettingEnrollmentProfile(ctx context.Context, id beta.DeviceManagementDepOnboardingSettingId, input beta.EnrollmentProfile, options CreateDepOnboardingSettingEnrollmentProfileOperationOptions) (result CreateDepOnboardingSettingEnrollmentProfileOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/enrollmentProfiles", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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
