package verificationdnsrecord

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

type ListVerificationDnsRecordsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DomainDnsRecord
}

type ListVerificationDnsRecordsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DomainDnsRecord
}

type ListVerificationDnsRecordsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVerificationDnsRecordsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVerificationDnsRecords ...
func (c VerificationDnsRecordClient) ListVerificationDnsRecords(ctx context.Context, id DomainId) (result ListVerificationDnsRecordsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVerificationDnsRecordsCustomPager{},
		Path:       fmt.Sprintf("%s/verificationDnsRecords", id.ID()),
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
		Values *[]stable.DomainDnsRecord `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVerificationDnsRecordsComplete retrieves all the results into a single object
func (c VerificationDnsRecordClient) ListVerificationDnsRecordsComplete(ctx context.Context, id DomainId) (ListVerificationDnsRecordsCompleteResult, error) {
	return c.ListVerificationDnsRecordsCompleteMatchingPredicate(ctx, id, DomainDnsRecordOperationPredicate{})
}

// ListVerificationDnsRecordsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VerificationDnsRecordClient) ListVerificationDnsRecordsCompleteMatchingPredicate(ctx context.Context, id DomainId, predicate DomainDnsRecordOperationPredicate) (result ListVerificationDnsRecordsCompleteResult, err error) {
	items := make([]stable.DomainDnsRecord, 0)

	resp, err := c.ListVerificationDnsRecords(ctx, id)
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

	result = ListVerificationDnsRecordsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
