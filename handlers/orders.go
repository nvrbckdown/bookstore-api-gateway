package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

// GetOrders forwards GET /orders request to order service
func GetOrders(orderServiceURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Pass along query parameters
		queryParams := c.Request.URL.Query().Encode()
		endpoint := "/api/v1/orders"
		if queryParams != "" {
			endpoint += "?" + queryParams
		}
		proxyRequest(c, orderServiceURL, endpoint, http.MethodGet)
	}
}

// GetOrder forwards GET /orders/:id request to order service
func GetOrder(orderServiceURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		proxyRequest(c, orderServiceURL, "/api/v1/orders/"+id, http.MethodGet)
	}
}

// CreateOrder forwards POST /orders request to order service
func CreateOrder(orderServiceURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxyRequest(c, orderServiceURL, "/api/v1/orders", http.MethodPost)
	}
}

// UpdateOrder forwards PUT /orders/:id request to order service
func UpdateOrder(orderServiceURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		proxyRequest(c, orderServiceURL, "/api/v1/orders/"+id, http.MethodPut)
	}
}

// DeleteOrder forwards DELETE /orders/:id request to order service
func DeleteOrder(orderServiceURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		proxyRequest(c, orderServiceURL, "/api/v1/orders/"+id, http.MethodDelete)
	}
}

// GetCustomerOrders forwards GET /orders/customer/:email request to order service
func GetCustomerOrders(orderServiceURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")
		// URL escape the email to handle special characters
		escapedEmail := url.PathEscape(email)
		proxyRequest(c, orderServiceURL, "/api/v1/orders/customer/"+escapedEmail, http.MethodGet)
	}
}

// proxyRequest forwards the request to the target service and returns the response
func proxyRequest(c *gin.Context, baseURL, path, method string) {
	// Create a new HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Read request body
	var requestBody []byte
	if c.Request.Body != nil {
		requestBody, _ = io.ReadAll(c.Request.Body)
		// Restore the request body for potential future middleware
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
	}

	// Create the request to the service
	req, err := http.NewRequest(method, baseURL+path, bytes.NewBuffer(requestBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create proxy request"})
		return
	}

	// Copy headers from original request
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Set content type if it's not already set
	if req.Header.Get("Content-Type") == "" && (method == http.MethodPost || method == http.MethodPut) {
		req.Header.Set("Content-Type", "application/json")
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service unavailable", "details": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Read response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read service response"})
		return
	}

	// Copy headers from the service response to our response
	for key, values := range resp.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}

	// Set response status code and body
	c.Status(resp.StatusCode)
	c.Writer.Write(responseBody)
}
