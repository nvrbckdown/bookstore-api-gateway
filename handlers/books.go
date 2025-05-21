package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks forwards GET /books request to bookstore service
func GetBooks(bookstoreURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxyRequest(c, bookstoreURL, "/api/books", http.MethodGet)
	}
}

// GetBook forwards GET /books/:id request to bookstore service
func GetBook(bookstoreURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		proxyRequest(c, bookstoreURL, "/api/books/"+id, http.MethodGet)
	}
}

// CreateBook forwards POST /books request to bookstore service
func CreateBook(bookstoreURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxyRequest(c, bookstoreURL, "/api/books", http.MethodPost)
	}
}

// UpdateBook forwards PUT /books/:id request to bookstore service
func UpdateBook(bookstoreURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		proxyRequest(c, bookstoreURL, "/api/books/"+id, http.MethodPut)
	}
}

// DeleteBook forwards DELETE /books/:id request to bookstore service
func DeleteBook(bookstoreURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		proxyRequest(c, bookstoreURL, "/api/books/"+id, http.MethodDelete)
	}
}

// GetAuthors forwards GET /authors request to bookstore service
func GetAuthors(bookstoreURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxyRequest(c, bookstoreURL, "/api/authors", http.MethodGet)
	}
}

// GetAuthor forwards GET /authors/:id request to bookstore service
func GetAuthor(bookstoreURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		proxyRequest(c, bookstoreURL, "/api/authors/"+id, http.MethodGet)
	}
}

// CreateAuthor forwards POST /authors request to bookstore service
func CreateAuthor(bookstoreURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxyRequest(c, bookstoreURL, "/api/authors", http.MethodPost)
	}
}

// UpdateAuthor forwards PUT /authors/:id request to bookstore service
func UpdateAuthor(bookstoreURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		proxyRequest(c, bookstoreURL, "/api/authors/"+id, http.MethodPut)
	}
}

// DeleteAuthor forwards DELETE /authors/:id request to bookstore service
func DeleteAuthor(bookstoreURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		proxyRequest(c, bookstoreURL, "/api/authors/"+id, http.MethodDelete)
	}
}

// GetAuthorBooks forwards GET /authors/:id/books request to bookstore service
func GetAuthorBooks(bookstoreURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		proxyRequest(c, bookstoreURL, "/api/authors/"+id+"/books", http.MethodGet)
	}
}
