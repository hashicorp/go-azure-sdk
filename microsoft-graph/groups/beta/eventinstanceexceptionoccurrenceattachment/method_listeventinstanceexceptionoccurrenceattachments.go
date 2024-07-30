package eventinstanceexceptionoccurrenceattachment

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

type ListEventInstanceExceptionOccurrenceAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Attachment
}

type ListEventInstanceExceptionOccurrenceAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Attachment
}

type ListEventInstanceExceptionOccurrenceAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEventInstanceExceptionOccurrenceAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEventInstanceExceptionOccurrenceAttachments ...
func (c EventInstanceExceptionOccurrenceAttachmentClient) ListEventInstanceExceptionOccurrenceAttachments(ctx context.Context, id GroupIdEventIdInstanceIdExceptionOccurrenceId) (result ListEventInstanceExceptionOccurrenceAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEventInstanceExceptionOccurrenceAttachmentsCustomPager{},
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
		Values *[]beta.Attachment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEventInstanceExceptionOccurrenceAttachmentsComplete retrieves all the results into a single object
func (c EventInstanceExceptionOccurrenceAttachmentClient) ListEventInstanceExceptionOccurrenceAttachmentsComplete(ctx context.Context, id GroupIdEventIdInstanceIdExceptionOccurrenceId) (ListEventInstanceExceptionOccurrenceAttachmentsCompleteResult, error) {
	return c.ListEventInstanceExceptionOccurrenceAttachmentsCompleteMatchingPredicate(ctx, id, AttachmentOperationPredicate{})
}

// ListEventInstanceExceptionOccurrenceAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EventInstanceExceptionOccurrenceAttachmentClient) ListEventInstanceExceptionOccurrenceAttachmentsCompleteMatchingPredicate(ctx context.Context, id GroupIdEventIdInstanceIdExceptionOccurrenceId, predicate AttachmentOperationPredicate) (result ListEventInstanceExceptionOccurrenceAttachmentsCompleteResult, err error) {
	items := make([]beta.Attachment, 0)

	resp, err := c.ListEventInstanceExceptionOccurrenceAttachments(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListEventInstanceExceptionOccurrenceAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
