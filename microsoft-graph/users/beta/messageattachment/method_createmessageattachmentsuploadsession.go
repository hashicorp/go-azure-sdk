package messageattachment

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

type CreateMessageAttachmentsUploadSessionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UploadSession
}

// CreateMessageAttachmentsUploadSession - Invoke action createUploadSession. Create an upload session that allows an
// app to iteratively upload ranges of a file, so as to attach the file to an Outlook item. The item can be a message or
// event. Use this approach to attach a file if the file size is between 3 MB and 150 MB. To attach a file that's
// smaller than 3 MB, do a POST operation on the attachments navigation property of the Outlook item; see how to do this
// for a message or for an event. As part of the response, this action returns an upload URL that you can use in
// subsequent sequential PUT queries. Request headers for each PUT operation let you specify the exact range of bytes to
// be uploaded. This allows transfer to be resumed, in case the network connection is dropped during upload. The
// following are the steps to attach a file to an Outlook item using an upload session: See attach large files to
// Outlook messages or events for an example.
func (c MessageAttachmentClient) CreateMessageAttachmentsUploadSession(ctx context.Context, id beta.UserIdMessageId, input CreateMessageAttachmentsUploadSessionRequest) (result CreateMessageAttachmentsUploadSessionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/attachments/createUploadSession", id.ID()),
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

	var model beta.UploadSession
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
