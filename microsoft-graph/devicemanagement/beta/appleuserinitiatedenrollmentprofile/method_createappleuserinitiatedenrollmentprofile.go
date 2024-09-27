package appleuserinitiatedenrollmentprofile

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAppleUserInitiatedEnrollmentProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AppleUserInitiatedEnrollmentProfile
}

type CreateAppleUserInitiatedEnrollmentProfileOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAppleUserInitiatedEnrollmentProfileOperationOptions() CreateAppleUserInitiatedEnrollmentProfileOperationOptions {
	return CreateAppleUserInitiatedEnrollmentProfileOperationOptions{}
}

func (o CreateAppleUserInitiatedEnrollmentProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAppleUserInitiatedEnrollmentProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAppleUserInitiatedEnrollmentProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAppleUserInitiatedEnrollmentProfile - Create new navigation property to appleUserInitiatedEnrollmentProfiles
// for deviceManagement
func (c AppleUserInitiatedEnrollmentProfileClient) CreateAppleUserInitiatedEnrollmentProfile(ctx context.Context, input beta.AppleUserInitiatedEnrollmentProfile, options CreateAppleUserInitiatedEnrollmentProfileOperationOptions) (result CreateAppleUserInitiatedEnrollmentProfileOperationResponse, err error) {
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
		Path:          "/deviceManagement/appleUserInitiatedEnrollmentProfiles",
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

	var model beta.AppleUserInitiatedEnrollmentProfile
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
