# Frequently Used HTTP Status Codes

## Top 10 Most Common in Backend Development

1. **200 OK** – Standard success response.
2. **201 Created** – Resource successfully created.
3. **204 No Content** – Successful request with no response body.
4. **400 Bad Request** – Invalid request or malformed input.
5. **401 Unauthorized** – Authentication required or failed.
6. **403 Forbidden** – Authenticated but not allowed.
7. **404 Not Found** – Requested resource not found.
8. **409 Conflict** – Resource conflict (e.g., duplicate entries).
9. **422 Unprocessable Entity** – Validation error.
10. **500 Internal Server Error** – Generic server-side error.

---

## Other Frequently Needed Status Codes

### 1xx — Informational

- **100 Continue** – Client can continue with request.

### 2xx — Success

- **202 Accepted** – Request accepted for async processing.

### 3xx — Redirection

- **301 Moved Permanently** – Resource moved; update client links.
- **302 Found** – Temporary redirect.
- **303 See Other** – Redirect using GET.
- **307 Temporary Redirect** – Method preserved during redirect.
- **308 Permanent Redirect** – Permanent redirect with method preserved.

### 4xx — Client Errors

- **405 Method Not Allowed** – Unsupported HTTP method.
- **410 Gone** – Resource removed permanently.
- **415 Unsupported Media Type** – Unsupported request format.
- **429 Too Many Requests** – Rate limit exceeded.

### 5xx — Server Errors

- **501 Not Implemented** – Endpoint or method not implemented.
- **502 Bad Gateway** – Upstream service returned an invalid response.
- **503 Service Unavailable** – Server is down or overloaded.
- **504 Gateway Timeout** – Upstream service timed out.
