package eventinstanceattachment

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

type ListEventInstanceAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Attachment
}

type ListEventInstanceAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Attachment
}

type ListEventInstanceAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEventInstanceAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEventInstanceAttachments ...
func (c EventInstanceAttachmentClient) ListEventInstanceAttachments(ctx context.Context, id MeEventIdInstanceId) (result ListEventInstanceAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEventInstanceAttachmentsCustomPager{},
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
		Values *[]stable.Attachment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEventInstanceAttachmentsComplete retrieves all the results into a single object
func (c EventInstanceAttachmentClient) ListEventInstanceAttachmentsComplete(ctx context.Context, id MeEventIdInstanceId) (ListEventInstanceAttachmentsCompleteResult, error) {
	return c.ListEventInstanceAttachmentsCompleteMatchingPredicate(ctx, id, AttachmentOperationPredicate{})
}

// ListEventInstanceAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EventInstanceAttachmentClient) ListEventInstanceAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeEventIdInstanceId, predicate AttachmentOperationPredicate) (result ListEventInstanceAttachmentsCompleteResult, err error) {
	items := make([]stable.Attachment, 0)

	resp, err := c.ListEventInstanceAttachments(ctx, id)
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

	result = ListEventInstanceAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
