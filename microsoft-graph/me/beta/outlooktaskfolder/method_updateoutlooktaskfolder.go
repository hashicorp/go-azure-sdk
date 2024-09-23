package outlooktaskfolder

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOutlookTaskFolderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateOutlookTaskFolderOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateOutlookTaskFolderOperationOptions() UpdateOutlookTaskFolderOperationOptions {
	return UpdateOutlookTaskFolderOperationOptions{}
}

func (o UpdateOutlookTaskFolderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateOutlookTaskFolderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateOutlookTaskFolderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateOutlookTaskFolder - Update outlooktaskfolder (deprecated). Update the writable properties of an Outlook task
// folder. You cannot change the name property value of the default task folder, 'Tasks'.
func (c OutlookTaskFolderClient) UpdateOutlookTaskFolder(ctx context.Context, id beta.MeOutlookTaskFolderId, input beta.OutlookTaskFolder, options UpdateOutlookTaskFolderOperationOptions) (result UpdateOutlookTaskFolderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
