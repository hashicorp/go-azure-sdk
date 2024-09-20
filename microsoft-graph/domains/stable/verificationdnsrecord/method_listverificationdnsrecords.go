package verificationdnsrecord

import (
	"context"
	"encoding/json"
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

type ListVerificationDnsRecordsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListVerificationDnsRecordsOperationOptions() ListVerificationDnsRecordsOperationOptions {
	return ListVerificationDnsRecordsOperationOptions{}
}

func (o ListVerificationDnsRecordsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListVerificationDnsRecordsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListVerificationDnsRecordsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListVerificationDnsRecords - List verificationDnsRecords. Retrieve a list of domainDnsRecord objects. You cannot use
// an associated domain with your Microsoft Entra tenant until ownership is verified. To verify the ownership of the
// domain, retrieve the domain verification records and add the details to the zone file of the domain. This can be done
// through the domain registrar or DNS server configuration. Root domains require verification. For example, contoso.com
// requires verification. If a root domain is verified, subdomains of the root domain are automatically verified. For
// example, subdomain.contoso.com is automatically be verified if contoso.com has been verified.
func (c VerificationDnsRecordClient) ListVerificationDnsRecords(ctx context.Context, id stable.DomainId, options ListVerificationDnsRecordsOperationOptions) (result ListVerificationDnsRecordsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListVerificationDnsRecordsCustomPager{},
		Path:          fmt.Sprintf("%s/verificationDnsRecords", id.ID()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.DomainDnsRecord, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalDomainDnsRecordImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.DomainDnsRecord (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListVerificationDnsRecordsComplete retrieves all the results into a single object
func (c VerificationDnsRecordClient) ListVerificationDnsRecordsComplete(ctx context.Context, id stable.DomainId, options ListVerificationDnsRecordsOperationOptions) (ListVerificationDnsRecordsCompleteResult, error) {
	return c.ListVerificationDnsRecordsCompleteMatchingPredicate(ctx, id, options, DomainDnsRecordOperationPredicate{})
}

// ListVerificationDnsRecordsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VerificationDnsRecordClient) ListVerificationDnsRecordsCompleteMatchingPredicate(ctx context.Context, id stable.DomainId, options ListVerificationDnsRecordsOperationOptions, predicate DomainDnsRecordOperationPredicate) (result ListVerificationDnsRecordsCompleteResult, err error) {
	items := make([]stable.DomainDnsRecord, 0)

	resp, err := c.ListVerificationDnsRecords(ctx, id, options)
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
