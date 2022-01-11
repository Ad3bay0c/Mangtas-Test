package handler

import (
	"github.com/Ad3bay0c/golang-testing/service"
	"github.com/gin-gonic/gin"
	"strings"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetMostUsedWords(c *gin.Context) {
	text := &struct {
		Text string `json:"text"`
	}{}
	if err := c.ShouldBindJSON(text); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if strings.TrimSpace(text.Text) == "" {
		c.JSON(400, gin.H{"error": "text is empty"})
		return
	}
	result := service.GetMostUsedWords(text.Text)

	c.JSON(200, result)
}

func (s *Server) SetUpRouter(r *gin.Engine) {
	r.POST("/", s.GetMostUsedWords)
}
