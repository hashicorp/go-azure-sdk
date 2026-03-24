package policyinvoke

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
)

const (
	// HeaderPolicyExternalEvaluations is the header used to pass the policy token to ARM.
	HeaderPolicyExternalEvaluations = "x-ms-policy-external-evaluations"
)

// PolicyInvokeOptions defines the options for the policy invoke middleware.
type PolicyInvokeOptions struct {
	// ChangeReference is the change reference associated with the operation.
	ChangeReference string
	// ApiVersion is the api-version of the acquirePolicyToken request.
	// If not specified, "2025-03-01" will be used.
	ApiVersion string
	// HTTPClient is the client used to make the policy API call.
	// If not specified, http.DefaultClient will be used.
	HTTPClient *http.Client
}

type acquirePolicyTokenRequest struct {
	Operation       operation `json:"operation"`
	ChangeReference string    `json:"changeReference,omitempty"`
}

type operation struct {
	Uri        string      `json:"uri"`
	HttpMethod string      `json:"httpMethod"`
	Content    interface{} `json:"content,omitempty"`
}

type acquirePolicyTokenResponse struct {
	Token string `json:"token"`
}

var subscriptionRegex = regexp.MustCompile(`(?i)/subscriptions/([^/]+)`)

// NewPolicyInvokeMiddleware returns a client.RequestMiddleware that acquires a policy token
// and attaches it to the request header.
func NewPolicyInvokeMiddleware(options PolicyInvokeOptions) client.RequestMiddleware {
	if options.ApiVersion == "" {
		options.ApiVersion = "2025-03-01" // Defaulting to this as it's the only supported version at time of writing
	}
	if options.HTTPClient == nil {
		options.HTTPClient = http.DefaultClient
	}

	return func(req *http.Request) (*http.Request, error) {
		// 1. Extract subscription ID to build the policy API URL
		matches := subscriptionRegex.FindStringSubmatch(req.URL.Path)
		if len(matches) < 2 {
			// If no subscription ID, we can't call the policy API.
			// For now we'll just continue as some requests might not have it.
			return req, nil
		}
		subscriptionId := matches[1]

		// 2. Read request body to include it in the policy request
		var body interface{}
		if req.Body != nil && req.Body != http.NoBody {
			reqBody, err := io.ReadAll(req.Body)
			if err != nil {
				return nil, fmt.Errorf("reading request body: %v", err)
			}
			// Reset body for the original request
			req.Body = io.NopCloser(bytes.NewBuffer(reqBody))

			// Try to unmarshal body to include it in the policy request
			var raw json.RawMessage
			if err := json.Unmarshal(reqBody, &raw); err == nil {
				body = raw
			} else {
				// If not JSON, send it as a string
				body = string(reqBody)
			}
		}

		// 3. Prepare policy API request
		policyReqBody := acquirePolicyTokenRequest{
			Operation: operation{
				Uri:        req.URL.String(),
				HttpMethod: req.Method,
				Content:    body,
			},
			ChangeReference: options.ChangeReference,
		}

		// Use the same host as the original request (e.g. management.azure.com)
		policyUrl := fmt.Sprintf("%s://%s/subscriptions/%s/providers/Microsoft.Authorization/acquirePolicyToken?api-version=%s",
			req.URL.Scheme, req.URL.Host, subscriptionId, options.ApiVersion)

		policyJson, err := json.Marshal(policyReqBody)
		if err != nil {
			return nil, fmt.Errorf("marshaling policy request: %v", err)
		}

		policyHttpRequest, err := http.NewRequestWithContext(req.Context(), http.MethodPost, policyUrl, bytes.NewReader(policyJson))
		if err != nil {
			return nil, fmt.Errorf("creating policy request: %v", err)
		}

		// Copy Authorization and other important headers from the original request
		policyHttpRequest.Header.Set("Authorization", req.Header.Get("Authorization"))
		policyHttpRequest.Header.Set("Content-Type", "application/json")
		if userAgent := req.Header.Get("User-Agent"); userAgent != "" {
			policyHttpRequest.Header.Set("User-Agent", userAgent)
		}

		// 4. Call policy API
		resp, err := options.HTTPClient.Do(policyHttpRequest)
		if err != nil {
			return nil, fmt.Errorf("calling policy API: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			// If policy API fails, we return an error to prevent the original request from proceeding without policy evaluation
			return nil, fmt.Errorf("policy API returned status %s", resp.Status)
		}

		var policyResp acquirePolicyTokenResponse
		if err := json.NewDecoder(resp.Body).Decode(&policyResp); err != nil {
			return nil, fmt.Errorf("decoding policy response: %v", err)
		}

		// 5. Attach token to header
		if policyResp.Token != "" {
			req.Header.Set(HeaderPolicyExternalEvaluations, policyResp.Token)
		}

		return req, nil
	}
}
