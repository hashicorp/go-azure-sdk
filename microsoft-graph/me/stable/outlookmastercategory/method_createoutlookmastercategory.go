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

type CreateOutlookMasterCategoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OutlookCategory
}

type CreateOutlookMasterCategoryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateOutlookMasterCategoryOperationOptions() CreateOutlookMasterCategoryOperationOptions {
	return CreateOutlookMasterCategoryOperationOptions{}
}

func (o CreateOutlookMasterCategoryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOutlookMasterCategoryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOutlookMasterCategoryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOutlookMasterCategory - Create Outlook category. Create an outlookCategory object in the user's master list of
// categories.
func (c OutlookMasterCategoryClient) CreateOutlookMasterCategory(ctx context.Context, input stable.OutlookCategory, options CreateOutlookMasterCategoryOperationOptions) (result CreateOutlookMasterCategoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/outlook/masterCategories",
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

	var model stable.OutlookCategory
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
