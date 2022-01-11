package handler

import (
	"encoding/json"
	"github.com/Ad3bay0c/golang-testing/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	router = gin.Default()
	s := NewServer()
	s.SetUpRouter(router)
	code := m.Run()
	os.Exit(code)
}
func TestApplicationHandler_GetMostUSedWords(t *testing.T) {
	response := httptest.NewRecorder()
	t.Run("Test for Invalid Request", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			t.Errorf("An Error Occurred: %v", err.Error())
		}
		router.ServeHTTP(response, req)
		assert.Contains(t, response.Body.String(), "error", "invalid request")
	})
	t.Run("Test for an Empty Input data", func(t *testing.T) {
		input := &struct {
			Text string `json:"text"`
		}{}
		text, _ := json.Marshal(input)
		req, err := http.NewRequest("POST", "/", strings.NewReader(string(text)))
		if err != nil {
			t.Errorf("An Error Occurred: %v", err.Error())
		}
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(response, req)

		assert.Contains(t, response.Body.String(), "text is empty")
	})

	tt := []struct {
		title string
		input *struct {
			Text string `json:"text"`
		}
		expected string
	}{
		{
			title: "Test for the highest occurring word",
			input: &struct {
				Text string `json:"text"`
			}{
				Text: "The quick brown fox jumps over the lazy dog",
			},
			expected: "{\"word\":\"the\",\"No_of_appearance\":2}",
		},
		{
			title: "Test for the second highest occurring word",
			input: &struct {
				Text string `json:"text"`
			}{
				Text: "The quick brown fox jumps over the lazy dog and the dog cried",
			},
			expected: "{\"word\":\"dog\",\"No_of_appearance\":2}",
		},
	}

	for _, tc := range tt {
		t.Run(tc.title, func(t *testing.T) {
			text, _ := json.Marshal(tc.input)
			req, err := http.NewRequest("POST", "/", strings.NewReader(string(text)))
			if err != nil {
				t.Errorf("An Error Occurred: %v", err.Error())
			}
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(response, req)

			assert.Contains(t, response.Body.String(), tc.expected)
		})
	}
}

func TestApplicationService_GetMostUsedWords(t *testing.T) {
	tt := []struct {
		title string
		input *struct {
			Text string `json:"text"`
		}
		expected int
		word     string
	}{
		{
			title: "Test for the highest occurring word",
			input: &struct {
				Text string `json:"text"`
			}{
				Text: "The quick brown fox jumps over the lazy dog and the dog cried",
			},
			expected: 3,
			word:     "the",
		},
		{
			title: "Test for the highest occurring word",
			input: &struct {
				Text string `json:"text"`
			}{
				Text: "The quick brown fox jumps over the lazy dog and the dog cried dog do dog",
			},
			expected: 4,
			word:     "dog",
		},
	}

	for _, tc := range tt {
		t.Run(tc.title, func(t *testing.T) {
			result := service.GetMostUsedWords(tc.input.Text)
			assert.Equal(t, tc.expected, result[0].Value)
		})
	}
}
