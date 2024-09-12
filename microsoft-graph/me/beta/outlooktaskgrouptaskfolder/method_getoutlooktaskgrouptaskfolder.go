package outlooktaskgrouptaskfolder

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetOutlookTaskGroupTaskFolderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.OutlookTaskFolder
}

type GetOutlookTaskGroupTaskFolderOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetOutlookTaskGroupTaskFolderOperationOptions() GetOutlookTaskGroupTaskFolderOperationOptions {
	return GetOutlookTaskGroupTaskFolderOperationOptions{}
}

func (o GetOutlookTaskGroupTaskFolderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOutlookTaskGroupTaskFolderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetOutlookTaskGroupTaskFolderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOutlookTaskGroupTaskFolder - Get taskFolders from me. The collection of task folders in the task group. Read-only.
// Nullable.
func (c OutlookTaskGroupTaskFolderClient) GetOutlookTaskGroupTaskFolder(ctx context.Context, id beta.MeOutlookTaskGroupIdTaskFolderId, options GetOutlookTaskGroupTaskFolderOperationOptions) (result GetOutlookTaskGroupTaskFolderOperationResponse, err error) {
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

	var model beta.OutlookTaskFolder
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
