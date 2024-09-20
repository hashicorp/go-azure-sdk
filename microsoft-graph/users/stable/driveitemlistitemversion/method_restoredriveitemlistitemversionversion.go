package driveitemlistitemversion

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreDriveItemListItemVersionVersionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RestoreDriveItemListItemVersionVersionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultRestoreDriveItemListItemVersionVersionOperationOptions() RestoreDriveItemListItemVersionVersionOperationOptions {
	return RestoreDriveItemListItemVersionVersionOperationOptions{}
}

func (o RestoreDriveItemListItemVersionVersionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RestoreDriveItemListItemVersionVersionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RestoreDriveItemListItemVersionVersionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RestoreDriveItemListItemVersionVersion - Invoke action restoreVersion. Restore a previous version of a ListItem to be
// the current version. This will create a new version with the contents of the previous version, but preserves all
// existing versions of the item.
func (c DriveItemListItemVersionClient) RestoreDriveItemListItemVersionVersion(ctx context.Context, id stable.UserIdDriveIdItemIdListItemVersionId, options RestoreDriveItemListItemVersionVersionOperationOptions) (result RestoreDriveItemListItemVersionVersionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/restoreVersion", id.ID()),
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
