// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package externalevaluation

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
)

const (
	// HeaderPolicyExternalEvaluations is the header used to pass the acquired policy token to ARM.
	HeaderPolicyExternalEvaluations = "x-ms-policy-external-evaluations"

	// HeaderForceSync ensures synchronous policy evaluation when acquiring a policy token.
	HeaderForceSync = "x-ms-force-sync"

	// DefaultApiVersion is the default api-version used for the acquirePolicyToken endpoint.
	DefaultApiVersion = "2025-11-01"

	// DisallowedByPolicyCode is the ARM error code returned when a request is rejected by Azure Policy.
	DisallowedByPolicyCode = "RequestDisallowedByPolicy"
)

type contextKey int

const (
	changeReferenceContextKey contextKey = iota
	retryAttemptedContextKey
)

// WithChangeReference returns a new context carrying the specified changeReference string
// for Azure Policy external evaluation token acquisition.
func WithChangeReference(ctx context.Context, changeRef string) context.Context {
	return context.WithValue(ctx, changeReferenceContextKey, changeRef)
}

// ChangeReferenceFromContext extracts the changeReference string from the context, if present.
func ChangeReferenceFromContext(ctx context.Context) (string, bool) {
	if ctx == nil {
		return "", false
	}
	v, ok := ctx.Value(changeReferenceContextKey).(string)
	return v, ok && v != ""
}

// Options defines the configuration options for the external evaluation policy middleware / transport.
type Options struct {
	// ChangeReference is an optional default change reference associated with operations
	// when one is not provided in the request's context.
	ChangeReference string

	// ApiVersion is the api-version of the acquirePolicyToken request.
	// If not specified, DefaultApiVersion ("2025-11-01") will be used.
	ApiVersion string

	// HTTPClient is the client used to make the acquirePolicyToken API call.
	// If not specified, http.DefaultClient will be used.
	HTTPClient *http.Client

	// Transport is the underlying http.RoundTripper used to execute ARM requests.
	// If not specified, http.DefaultTransport will be used.
	Transport http.RoundTripper
}

// ExternalEvaluationOptions is an alias for Options.
type ExternalEvaluationOptions = Options

// AcquirePolicyTokenRequest represents the payload sent to the Microsoft.Authorization/acquirePolicyToken API.
type AcquirePolicyTokenRequest struct {
	Operation       AcquirePolicyTokenOperation `json:"operation"`
	ChangeReference string                      `json:"changeReference,omitempty"`
}

// AcquirePolicyTokenOperation describes the target ARM operation being evaluated.
type AcquirePolicyTokenOperation struct {
	Uri        string          `json:"uri"`
	HttpMethod string          `json:"httpMethod"`
	Content    json.RawMessage `json:"content,omitempty"`
}

// AcquirePolicyTokenResponse represents the response returned by Microsoft.Authorization/acquirePolicyToken.
type AcquirePolicyTokenResponse struct {
	Result     string         `json:"result,omitempty"`
	Token      string         `json:"token,omitempty"`
	TokenId    string         `json:"tokenId,omitempty"`
	Expiration string         `json:"expiration,omitempty"`
	Results    []PolicyResult `json:"results,omitempty"`
}

// PolicyResult represents an individual policy evaluation outcome within AcquirePolicyTokenResponse.
type PolicyResult struct {
	PolicyAssignmentId      string                   `json:"policyAssignmentId,omitempty"`
	PolicyDefinitionId      string                   `json:"policyDefinitionId,omitempty"`
	Result                  string                   `json:"result,omitempty"`
	Message                 string                   `json:"message,omitempty"`
	PolicyEvaluationDetails *PolicyEvaluationDetails `json:"policyEvaluationDetails,omitempty"`
}

// PolicyEvaluationDetails contains detailed evaluation information for a PolicyResult.
type PolicyEvaluationDetails struct {
	Reason string `json:"reason,omitempty"`
}

// MissingTokenError models the HTTP 403 Forbidden ARM error payload for RequestDisallowedByPolicy.
type MissingTokenError struct {
	Error *MissingTokenErrorBody `json:"error,omitempty"`
}

// MissingTokenErrorBody represents the inner error structure of an ARM policy denial response.
type MissingTokenErrorBody struct {
	Code           string                       `json:"code,omitempty"`
	Message        string                       `json:"message,omitempty"`
	AdditionalInfo []MissingTokenAdditionalInfo `json:"additionalInfo,omitempty"`
	Details        []MissingTokenErrorBody      `json:"details,omitempty"`
}

// MissingTokenAdditionalInfo represents an entry in error.additionalInfo.
type MissingTokenAdditionalInfo struct {
	Type string          `json:"type,omitempty"`
	Info json.RawMessage `json:"info,omitempty"`
}

// MissingTokenInfo represents the info object within an error.additionalInfo entry.
type MissingTokenInfo struct {
	EvaluationDetails *MissingTokenEvaluationDetails `json:"evaluationDetails,omitempty"`
}

// MissingTokenEvaluationDetails holds missingPolicyTokenDetails inside evaluationDetails.
type MissingTokenEvaluationDetails struct {
	MissingPolicyTokenDetails []json.RawMessage `json:"missingPolicyTokenDetails,omitempty"`
}

// PolicyEvaluationError is returned when acquirePolicyToken returns HTTP 200 OK without a token,
// indicating that external policy evaluation denied or failed the operation.
type PolicyEvaluationError struct {
	Result  string
	Results []PolicyResult
	Reasons []string
}

func (e *PolicyEvaluationError) Error() string {
	if len(e.Reasons) > 0 {
		return fmt.Sprintf("external policy evaluation failed (result: %q): %s", e.Result, strings.Join(e.Reasons, "; "))
	}
	return fmt.Sprintf("external policy evaluation failed (result: %q): no policy token returned", e.Result)
}

var (
	subscriptionScopeRegex    = regexp.MustCompile(`(?i)/subscriptions/[^/]+`)
	managementGroupScopeRegex = regexp.MustCompile(`(?i)/providers/Microsoft\.Management/managementGroups/[^/]+`)
)

// extractPolicyTokenScope extracts the ARM scope prefix for the acquirePolicyToken endpoint
// from the request URL path. Both Subscription (/subscriptions/{subscriptionId}) and
// Management Group (/providers/Microsoft.Management/managementGroups/{managementGroupId}) scopes are supported.
func extractPolicyTokenScope(path string) (string, bool) {
	subLoc := subscriptionScopeRegex.FindStringIndex(path)
	mgLoc := managementGroupScopeRegex.FindStringIndex(path)

	switch {
	case subLoc != nil && mgLoc != nil:
		if subLoc[0] <= mgLoc[0] {
			return path[subLoc[0]:subLoc[1]], true
		}
		return path[mgLoc[0]:mgLoc[1]], true
	case subLoc != nil:
		return path[subLoc[0]:subLoc[1]], true
	case mgLoc != nil:
		return path[mgLoc[0]:mgLoc[1]], true
	default:
		return "", false
	}
}

// marshalAcquirePolicyTokenRequest serializes AcquirePolicyTokenRequest into JSON while preserving
// req.Operation.Content byte-for-byte.
//
// IMPORTANT: Do NOT replace this function with standard encoding/json.Marshal.
// Go's encoding/json.Marshal invokes json.Compact on all struct fields—including json.RawMessage
// and custom json.Marshaler outputs—which strips whitespace and escapes HTML characters ('<', '>', '&').
// Because the token returned by acquirePolicyToken is cryptographically bound to the target request body
// via an internal contentHash, any whitespace compaction or re-serialization discrepancy between
// operation.content and the target ARM request body will cause ARM to reject the token on retry.
func marshalAcquirePolicyTokenRequest(req AcquirePolicyTokenRequest) ([]byte, error) {
	uriJSON, err := marshalJSONStringNoEscape(req.Operation.Uri)
	if err != nil {
		return nil, fmt.Errorf("marshaling operation.uri: %w", err)
	}
	methodJSON, err := marshalJSONStringNoEscape(req.Operation.HttpMethod)
	if err != nil {
		return nil, fmt.Errorf("marshaling operation.httpMethod: %w", err)
	}

	var buf bytes.Buffer
	buf.WriteString(`{"operation":{"uri":`)
	buf.Write(uriJSON)
	buf.WriteString(`,"httpMethod":`)
	buf.Write(methodJSON)

	if len(req.Operation.Content) > 0 {
		buf.WriteString(`,"content":`)
		buf.Write(req.Operation.Content)
	}

	buf.WriteByte('}')

	if req.ChangeReference != "" {
		changeRefJSON, err := marshalJSONStringNoEscape(req.ChangeReference)
		if err != nil {
			return nil, fmt.Errorf("marshaling changeReference: %w", err)
		}
		buf.WriteString(`,"changeReference":`)
		buf.Write(changeRefJSON)
	}

	buf.WriteByte('}')
	return buf.Bytes(), nil
}

func marshalJSONStringNoEscape(s string) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(s); err != nil {
		return nil, err
	}
	// json.Encoder.Encode appends a trailing newline; strip it.
	return bytes.TrimSuffix(buf.Bytes(), []byte("\n")), nil
}

type roundTripper struct {
	options Options
}

// NewRoundTripper returns an http.RoundTripper that reactively intercepts HTTP 403 Forbidden
// RequestDisallowedByPolicy responses containing missingPolicyTokenDetails, acquires a policy token
// from Microsoft.Authorization/acquirePolicyToken, and retries the request once with the token attached.
func NewRoundTripper(options Options) http.RoundTripper {
	return &roundTripper{
		options: normalizeOptions(options),
	}
}

// NewMiddleware returns a client.ResponseMiddleware that reactively intercepts HTTP 403 Forbidden
// RequestDisallowedByPolicy responses containing missingPolicyTokenDetails, acquires a policy token
// from Microsoft.Authorization/acquirePolicyToken, and retries the request once with the token attached.
func NewMiddleware(options Options) client.ResponseMiddleware {
	opts := normalizeOptions(options)
	rt := &roundTripper{options: opts}
	return func(req *http.Request, resp *http.Response) (*http.Response, error) {
		if req == nil || resp == nil {
			return resp, nil
		}
		if !isMutatingMethod(req.Method) || alreadyAttemptedRetry(req) {
			return resp, nil
		}
		var bodyBytes []byte
		if req.Method != http.MethodDelete {
			var err error
			bodyBytes, err = readOrGetRequestBody(req)
			if err != nil {
				return nil, err
			}
		}
		return rt.handleResponse(req, resp, bodyBytes)
	}
}

func normalizeOptions(options Options) Options {
	if options.ApiVersion == "" {
		options.ApiVersion = DefaultApiVersion
	}
	if options.Transport == nil {
		options.Transport = http.DefaultTransport
	}
	if options.HTTPClient == nil {
		options.HTTPClient = &http.Client{
			Transport: options.Transport,
		}
	}
	return options
}

func (rt *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if !isMutatingMethod(req.Method) || alreadyAttemptedRetry(req) {
		return rt.options.Transport.RoundTrip(req)
	}

	// Safely buffer the request body before dispatching so that it can be reused
	// byte-for-byte in acquirePolicyToken and during the guarded retry.
	bodyBytes, err := bufferRequestBody(req)
	if err != nil {
		return nil, err
	}

	resp, err := rt.options.Transport.RoundTrip(req)
	if err != nil || resp == nil {
		return resp, err
	}

	return rt.handleResponse(req, resp, bodyBytes)
}

func (rt *roundTripper) handleResponse(req *http.Request, resp *http.Response, bodyBytes []byte) (*http.Response, error) {
	if resp.StatusCode != http.StatusForbidden {
		return resp, nil
	}

	scopePrefix, ok := extractPolicyTokenScope(req.URL.Path)
	if !ok {
		return resp, nil
	}

	isMissingToken, err := inspectMissingPolicyTokenError(resp)
	if err != nil || !isMissingToken {
		return resp, err
	}

	token, err := rt.acquirePolicyToken(req, scopePrefix, bodyBytes)
	if err != nil {
		_ = resp.Body.Close()
		return nil, err
	}

	// Close the initial 403 response body before retrying.
	_ = resp.Body.Close()

	// Clone the original request, mark the context to prevent infinite retry loops,
	// attach the verbatim PoP token header, and restore the exact request body bytes.
	retryCtx := context.WithValue(req.Context(), retryAttemptedContextKey, true)
	retryReq := req.Clone(retryCtx)
	retryReq.Header.Set(HeaderPolicyExternalEvaluations, token)

	if bodyBytes != nil {
		retryReq.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		retryReq.GetBody = func() (io.ReadCloser, error) {
			return io.NopCloser(bytes.NewReader(bodyBytes)), nil
		}
		retryReq.ContentLength = int64(len(bodyBytes))
	} else {
		retryReq.Body = nil
		retryReq.GetBody = nil
		retryReq.ContentLength = 0
	}

	return rt.options.Transport.RoundTrip(retryReq)
}

func (rt *roundTripper) acquirePolicyToken(req *http.Request, scopePrefix string, bodyBytes []byte) (string, error) {
	changeRef, ok := ChangeReferenceFromContext(req.Context())
	if !ok {
		changeRef = rt.options.ChangeReference
	}

	op := AcquirePolicyTokenOperation{
		Uri:        req.URL.RequestURI(),
		HttpMethod: req.Method,
	}
	if !strings.EqualFold(req.Method, http.MethodDelete) && len(bodyBytes) > 0 {
		op.Content = json.RawMessage(bodyBytes)
	}

	tokenReq := AcquirePolicyTokenRequest{
		Operation:       op,
		ChangeReference: changeRef,
	}

	payload, err := marshalAcquirePolicyTokenRequest(tokenReq)
	if err != nil {
		return "", fmt.Errorf("marshaling acquirePolicyToken request: %w", err)
	}

	policyURL := fmt.Sprintf("%s://%s%s/providers/Microsoft.Authorization/acquirePolicyToken?api-version=%s",
		req.URL.Scheme, req.URL.Host, scopePrefix, rt.options.ApiVersion)

	httpReq, err := http.NewRequestWithContext(req.Context(), http.MethodPost, policyURL, bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("creating acquirePolicyToken request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set(HeaderForceSync, "true")
	if authHeader := req.Header.Get("Authorization"); authHeader != "" {
		httpReq.Header.Set("Authorization", authHeader)
	}
	if userAgent := req.Header.Get("User-Agent"); userAgent != "" {
		httpReq.Header.Set("User-Agent", userAgent)
	}
	if corrID := req.Header.Get("X-Ms-Correlation-Request-Id"); corrID != "" {
		httpReq.Header.Set("X-Ms-Correlation-Request-Id", corrID)
	}

	resp, err := rt.options.HTTPClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("calling acquirePolicyToken API: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading acquirePolicyToken response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		retryAfterMsg := ""
		if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
			if err := waitRetryAfter(req.Context(), retryAfter); err != nil {
				return "", err
			}
			retryAfterMsg = fmt.Sprintf(" (retry-after: %s)", retryAfter)
		}
		if len(respBytes) > 0 {
			return "", fmt.Errorf("acquirePolicyToken API returned status %s%s: %s", resp.Status, retryAfterMsg, string(respBytes))
		}
		return "", fmt.Errorf("acquirePolicyToken API returned status %s%s", resp.Status, retryAfterMsg)
	}

	var tokenResp AcquirePolicyTokenResponse
	if err := json.Unmarshal(respBytes, &tokenResp); err != nil {
		return "", fmt.Errorf("decoding acquirePolicyToken response: %w", err)
	}

	if tokenResp.Token == "" {
		var reasons []string
		for _, r := range tokenResp.Results {
			if r.Message != "" {
				reasons = append(reasons, r.Message)
			}
			if r.PolicyEvaluationDetails != nil && r.PolicyEvaluationDetails.Reason != "" {
				reasons = append(reasons, r.PolicyEvaluationDetails.Reason)
			}
		}
		return "", &PolicyEvaluationError{
			Result:  tokenResp.Result,
			Results: tokenResp.Results,
			Reasons: reasons,
		}
	}

	return tokenResp.Token, nil
}

func waitRetryAfter(ctx context.Context, retryAfter string) error {
	seconds, err := strconv.Atoi(strings.TrimSpace(retryAfter))
	if err != nil || seconds <= 0 {
		return nil
	}
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func isMutatingMethod(method string) bool {
	switch strings.ToUpper(method) {
	case http.MethodPut, http.MethodPost, http.MethodPatch, http.MethodDelete:
		return true
	default:
		return false
	}
}

func alreadyAttemptedRetry(req *http.Request) bool {
	if req.Header.Get(HeaderPolicyExternalEvaluations) != "" {
		return true
	}
	if req.Context() != nil {
		if attempted, ok := req.Context().Value(retryAttemptedContextKey).(bool); ok && attempted {
			return true
		}
	}
	return false
}

func bufferRequestBody(req *http.Request) ([]byte, error) {
	if req.Body == nil || req.Body == http.NoBody {
		return nil, nil
	}
	bodyBytes, err := io.ReadAll(req.Body)
	if closeErr := req.Body.Close(); closeErr != nil && err == nil {
		err = closeErr
	}
	if err != nil {
		return nil, fmt.Errorf("reading request body: %w", err)
	}
	req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(bodyBytes)), nil
	}
	req.ContentLength = int64(len(bodyBytes))
	return bodyBytes, nil
}

func readOrGetRequestBody(req *http.Request) ([]byte, error) {
	if req.GetBody != nil {
		rc, err := req.GetBody()
		if err != nil {
			return nil, fmt.Errorf("getting request body copy: %w", err)
		}
		defer rc.Close()
		b, err := io.ReadAll(rc)
		if err != nil {
			return nil, fmt.Errorf("reading request body copy: %w", err)
		}
		return b, nil
	}
	return bufferRequestBody(req)
}

func inspectMissingPolicyTokenError(resp *http.Response) (bool, error) {
	if resp == nil || resp.Body == nil || resp.Body == http.NoBody {
		return false, nil
	}

	respBytes, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		return false, fmt.Errorf("reading 403 response body: %w", err)
	}
	// Always restore the response body so downstream callers can read it if not intercepted.
	resp.Body = io.NopCloser(bytes.NewReader(respBytes))

	if len(respBytes) == 0 {
		return false, nil
	}

	var parsed MissingTokenError
	if err := json.Unmarshal(respBytes, &parsed); err != nil || parsed.Error == nil {
		return false, nil
	}

	return hasMissingPolicyTokenDetails(parsed.Error), nil
}

func hasMissingPolicyTokenDetails(errBody *MissingTokenErrorBody) bool {
	if errBody == nil {
		return false
	}

	if strings.EqualFold(errBody.Code, DisallowedByPolicyCode) {
		for _, addInfo := range errBody.AdditionalInfo {
			if len(addInfo.Info) == 0 {
				continue
			}
			var info MissingTokenInfo
			if err := json.Unmarshal(addInfo.Info, &info); err == nil {
				if info.EvaluationDetails != nil && len(info.EvaluationDetails.MissingPolicyTokenDetails) > 0 {
					return true
				}
			}
		}
	}

	for i := range errBody.Details {
		if hasMissingPolicyTokenDetails(&errBody.Details[i]) {
			return true
		}
	}

	return false
}
