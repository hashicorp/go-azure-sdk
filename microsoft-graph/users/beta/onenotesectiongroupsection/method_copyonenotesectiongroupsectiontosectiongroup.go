package onenotesectiongroupsection

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

type CopyOnenoteSectionGroupSectionToSectionGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.OnenoteOperation
}

type CopyOnenoteSectionGroupSectionToSectionGroupOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCopyOnenoteSectionGroupSectionToSectionGroupOperationOptions() CopyOnenoteSectionGroupSectionToSectionGroupOperationOptions {
	return CopyOnenoteSectionGroupSectionToSectionGroupOperationOptions{}
}

func (o CopyOnenoteSectionGroupSectionToSectionGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CopyOnenoteSectionGroupSectionToSectionGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CopyOnenoteSectionGroupSectionToSectionGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CopyOnenoteSectionGroupSectionToSectionGroup - Invoke action copyToSectionGroup. Copies a section to a specific
// section group. For Copy operations, you follow an asynchronous calling pattern: First call the Copy action, and then
// poll the operation endpoint for the result.
func (c OnenoteSectionGroupSectionClient) CopyOnenoteSectionGroupSectionToSectionGroup(ctx context.Context, id beta.UserIdOnenoteSectionGroupIdSectionId, input CopyOnenoteSectionGroupSectionToSectionGroupRequest, options CopyOnenoteSectionGroupSectionToSectionGroupOperationOptions) (result CopyOnenoteSectionGroupSectionToSectionGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/copyToSectionGroup", id.ID()),
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

	var model beta.OnenoteOperation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
