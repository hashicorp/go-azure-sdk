package contactfoldercontactphoto

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

type RemoveContactFolderContactPhotoValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveContactFolderContactPhotoValueOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveContactFolderContactPhotoValueOperationOptions() RemoveContactFolderContactPhotoValueOperationOptions {
	return RemoveContactFolderContactPhotoValueOperationOptions{}
}

func (o RemoveContactFolderContactPhotoValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveContactFolderContactPhotoValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveContactFolderContactPhotoValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveContactFolderContactPhotoValue - Delete media content for the navigation property photo in me. Optional contact
// picture. You can get or set a photo for a contact.
func (c ContactFolderContactPhotoClient) RemoveContactFolderContactPhotoValue(ctx context.Context, id stable.MeContactFolderIdContactId, options RemoveContactFolderContactPhotoValueOperationOptions) (result RemoveContactFolderContactPhotoValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/photo/$value", id.ID()),
		RetryFunc:     options.RetryFunc,
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
