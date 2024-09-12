package intunebrandingprofileassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetIntuneBrandingProfileAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.IntuneBrandingProfileAssignment
}

type GetIntuneBrandingProfileAssignmentOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetIntuneBrandingProfileAssignmentOperationOptions() GetIntuneBrandingProfileAssignmentOperationOptions {
	return GetIntuneBrandingProfileAssignmentOperationOptions{}
}

func (o GetIntuneBrandingProfileAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetIntuneBrandingProfileAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetIntuneBrandingProfileAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetIntuneBrandingProfileAssignment - Get assignments from deviceManagement. The list of group assignments for the
// branding profile
func (c IntuneBrandingProfileAssignmentClient) GetIntuneBrandingProfileAssignment(ctx context.Context, id beta.DeviceManagementIntuneBrandingProfileIdAssignmentId, options GetIntuneBrandingProfileAssignmentOperationOptions) (result GetIntuneBrandingProfileAssignmentOperationResponse, err error) {
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

	var model beta.IntuneBrandingProfileAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
