package meoutlooktaskfoldertaskattachment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeOutlookTaskFolderByIdTaskByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListMeOutlookTaskFolderByIdTaskByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListMeOutlookTaskFolderByIdTaskByIdAttachments ...
func (c MeOutlookTaskFolderTaskAttachmentClient) ListMeOutlookTaskFolderByIdTaskByIdAttachments(ctx context.Context, id MeOutlookTaskFolderTaskId) (result ListMeOutlookTaskFolderByIdTaskByIdAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/attachments", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]models.AttachmentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeOutlookTaskFolderByIdTaskByIdAttachmentsComplete retrieves all the results into a single object
func (c MeOutlookTaskFolderTaskAttachmentClient) ListMeOutlookTaskFolderByIdTaskByIdAttachmentsComplete(ctx context.Context, id MeOutlookTaskFolderTaskId) (ListMeOutlookTaskFolderByIdTaskByIdAttachmentsCompleteResult, error) {
	return c.ListMeOutlookTaskFolderByIdTaskByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListMeOutlookTaskFolderByIdTaskByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOutlookTaskFolderTaskAttachmentClient) ListMeOutlookTaskFolderByIdTaskByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeOutlookTaskFolderTaskId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListMeOutlookTaskFolderByIdTaskByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListMeOutlookTaskFolderByIdTaskByIdAttachments(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListMeOutlookTaskFolderByIdTaskByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
