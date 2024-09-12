package windowsfeatureupdateprofile

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

type DeleteWindowsFeatureUpdateProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteWindowsFeatureUpdateProfileOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteWindowsFeatureUpdateProfileOperationOptions() DeleteWindowsFeatureUpdateProfileOperationOptions {
	return DeleteWindowsFeatureUpdateProfileOperationOptions{}
}

func (o DeleteWindowsFeatureUpdateProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteWindowsFeatureUpdateProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteWindowsFeatureUpdateProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteWindowsFeatureUpdateProfile - Delete navigation property windowsFeatureUpdateProfiles for deviceManagement
func (c WindowsFeatureUpdateProfileClient) DeleteWindowsFeatureUpdateProfile(ctx context.Context, id beta.DeviceManagementWindowsFeatureUpdateProfileId, options DeleteWindowsFeatureUpdateProfileOperationOptions) (result DeleteWindowsFeatureUpdateProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
