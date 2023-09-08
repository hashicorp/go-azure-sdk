package usereventinstanceexceptionoccurrenceattachment

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

type ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments ...
func (c UserEventInstanceExceptionOccurrenceAttachmentClient) ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx context.Context, id UserEventInstanceExceptionOccurrenceId) (result ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsOperationResponse, err error) {
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

// ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete retrieves all the results into a single object
func (c UserEventInstanceExceptionOccurrenceAttachmentClient) ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsComplete(ctx context.Context, id UserEventInstanceExceptionOccurrenceId) (ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteResult, error) {
	return c.ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserEventInstanceExceptionOccurrenceAttachmentClient) ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id UserEventInstanceExceptionOccurrenceId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachments(ctx, id)
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

	result = ListUserByIdEventByIdInstanceByIdExceptionOccurrenceByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
