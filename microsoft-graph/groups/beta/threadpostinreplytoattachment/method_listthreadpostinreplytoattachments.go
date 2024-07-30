package threadpostinreplytoattachment

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

type ListThreadPostInReplyToAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Attachment
}

type ListThreadPostInReplyToAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Attachment
}

type ListThreadPostInReplyToAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListThreadPostInReplyToAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListThreadPostInReplyToAttachments ...
func (c ThreadPostInReplyToAttachmentClient) ListThreadPostInReplyToAttachments(ctx context.Context, id GroupIdThreadIdPostId) (result ListThreadPostInReplyToAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListThreadPostInReplyToAttachmentsCustomPager{},
		Path:       fmt.Sprintf("%s/inReplyTo/attachments", id.ID()),
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

// ListThreadPostInReplyToAttachmentsComplete retrieves all the results into a single object
func (c ThreadPostInReplyToAttachmentClient) ListThreadPostInReplyToAttachmentsComplete(ctx context.Context, id GroupIdThreadIdPostId) (ListThreadPostInReplyToAttachmentsCompleteResult, error) {
	return c.ListThreadPostInReplyToAttachmentsCompleteMatchingPredicate(ctx, id, AttachmentOperationPredicate{})
}

// ListThreadPostInReplyToAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ThreadPostInReplyToAttachmentClient) ListThreadPostInReplyToAttachmentsCompleteMatchingPredicate(ctx context.Context, id GroupIdThreadIdPostId, predicate AttachmentOperationPredicate) (result ListThreadPostInReplyToAttachmentsCompleteResult, err error) {
	items := make([]beta.Attachment, 0)

	resp, err := c.ListThreadPostInReplyToAttachments(ctx, id)
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

	result = ListThreadPostInReplyToAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
