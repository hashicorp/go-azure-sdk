package mecalendareventattachment

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

type ListMeCalendarByIdEventByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListMeCalendarByIdEventByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListMeCalendarByIdEventByIdAttachments ...
func (c MeCalendarEventAttachmentClient) ListMeCalendarByIdEventByIdAttachments(ctx context.Context, id MeCalendarEventId) (result ListMeCalendarByIdEventByIdAttachmentsOperationResponse, err error) {
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

// ListMeCalendarByIdEventByIdAttachmentsComplete retrieves all the results into a single object
func (c MeCalendarEventAttachmentClient) ListMeCalendarByIdEventByIdAttachmentsComplete(ctx context.Context, id MeCalendarEventId) (ListMeCalendarByIdEventByIdAttachmentsCompleteResult, error) {
	return c.ListMeCalendarByIdEventByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListMeCalendarByIdEventByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarEventAttachmentClient) ListMeCalendarByIdEventByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeCalendarEventId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListMeCalendarByIdEventByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListMeCalendarByIdEventByIdAttachments(ctx, id)
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

	result = ListMeCalendarByIdEventByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
