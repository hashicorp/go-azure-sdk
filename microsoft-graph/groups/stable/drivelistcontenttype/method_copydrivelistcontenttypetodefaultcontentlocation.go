package drivelistcontenttype

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

type CopyDriveListContentTypeToDefaultContentLocationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CopyDriveListContentTypeToDefaultContentLocation - Invoke action copyToDefaultContentLocation. Copy a file to a
// default content location in a content type. The file can then be added as a default file or template via a POST
// operation.
func (c DriveListContentTypeClient) CopyDriveListContentTypeToDefaultContentLocation(ctx context.Context, id stable.GroupIdDriveIdListContentTypeId, input CopyDriveListContentTypeToDefaultContentLocationRequest) (result CopyDriveListContentTypeToDefaultContentLocationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/copyToDefaultContentLocation", id.ID()),
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
