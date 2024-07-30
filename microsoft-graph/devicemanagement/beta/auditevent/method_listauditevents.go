package auditevent

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

type ListAuditEventsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AuditEvent
}

type ListAuditEventsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AuditEvent
}

type ListAuditEventsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuditEventsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuditEvents ...
func (c AuditEventClient) ListAuditEvents(ctx context.Context) (result ListAuditEventsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAuditEventsCustomPager{},
		Path:       "/deviceManagement/auditEvents",
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
		Values *[]beta.AuditEvent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAuditEventsComplete retrieves all the results into a single object
func (c AuditEventClient) ListAuditEventsComplete(ctx context.Context) (ListAuditEventsCompleteResult, error) {
	return c.ListAuditEventsCompleteMatchingPredicate(ctx, AuditEventOperationPredicate{})
}

// ListAuditEventsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuditEventClient) ListAuditEventsCompleteMatchingPredicate(ctx context.Context, predicate AuditEventOperationPredicate) (result ListAuditEventsCompleteResult, err error) {
	items := make([]beta.AuditEvent, 0)

	resp, err := c.ListAuditEvents(ctx)
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

	result = ListAuditEventsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
