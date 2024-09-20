package userexperienceanalyticsimpactingprocess

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateUserExperienceAnalyticsImpactingProcessOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateUserExperienceAnalyticsImpactingProcessOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateUserExperienceAnalyticsImpactingProcessOperationOptions() UpdateUserExperienceAnalyticsImpactingProcessOperationOptions {
	return UpdateUserExperienceAnalyticsImpactingProcessOperationOptions{}
}

func (o UpdateUserExperienceAnalyticsImpactingProcessOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateUserExperienceAnalyticsImpactingProcessOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateUserExperienceAnalyticsImpactingProcessOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateUserExperienceAnalyticsImpactingProcess - Update the navigation property
// userExperienceAnalyticsImpactingProcess in deviceManagement
func (c UserExperienceAnalyticsImpactingProcessClient) UpdateUserExperienceAnalyticsImpactingProcess(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsImpactingProcessId, input beta.UserExperienceAnalyticsImpactingProcess, options UpdateUserExperienceAnalyticsImpactingProcessOperationOptions) (result UpdateUserExperienceAnalyticsImpactingProcessOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
