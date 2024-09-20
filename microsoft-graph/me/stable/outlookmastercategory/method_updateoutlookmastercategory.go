package outlookmastercategory

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOutlookMasterCategoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateOutlookMasterCategoryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateOutlookMasterCategoryOperationOptions() UpdateOutlookMasterCategoryOperationOptions {
	return UpdateOutlookMasterCategoryOperationOptions{}
}

func (o UpdateOutlookMasterCategoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateOutlookMasterCategoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateOutlookMasterCategoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateOutlookMasterCategory - Update outlookCategory. Update the writable property, color, of the specified
// outlookCategory object. You can't modify the displayName property once you have created the category.
func (c OutlookMasterCategoryClient) UpdateOutlookMasterCategory(ctx context.Context, id stable.MeOutlookMasterCategoryId, input stable.OutlookCategory, options UpdateOutlookMasterCategoryOperationOptions) (result UpdateOutlookMasterCategoryOperationResponse, err error) {
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
