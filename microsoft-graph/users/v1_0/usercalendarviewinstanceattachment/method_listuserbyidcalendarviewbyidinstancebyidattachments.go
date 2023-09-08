package usercalendarviewinstanceattachment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdCalendarViewByIdInstanceByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListUserByIdCalendarViewByIdInstanceByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListUserByIdCalendarViewByIdInstanceByIdAttachments ...
func (c UserCalendarViewInstanceAttachmentClient) ListUserByIdCalendarViewByIdInstanceByIdAttachments(ctx context.Context, id UserCalendarViewInstanceId) (result ListUserByIdCalendarViewByIdInstanceByIdAttachmentsOperationResponse, err error) {
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

// ListUserByIdCalendarViewByIdInstanceByIdAttachmentsComplete retrieves all the results into a single object
func (c UserCalendarViewInstanceAttachmentClient) ListUserByIdCalendarViewByIdInstanceByIdAttachmentsComplete(ctx context.Context, id UserCalendarViewInstanceId) (ListUserByIdCalendarViewByIdInstanceByIdAttachmentsCompleteResult, error) {
	return c.ListUserByIdCalendarViewByIdInstanceByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarViewByIdInstanceByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarViewInstanceAttachmentClient) ListUserByIdCalendarViewByIdInstanceByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id UserCalendarViewInstanceId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListUserByIdCalendarViewByIdInstanceByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarViewByIdInstanceByIdAttachments(ctx, id)
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

	result = ListUserByIdCalendarViewByIdInstanceByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
