package server

import (
	"github.com/TheoEb/scalingo_test/backend/src/models"
	"log"
	"net/http"
	"strings"

	"github.com/TheoEb/scalingo_test/backend/src/github"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine       *gin.Engine
	GithubClient *github.Client
	Port         string
}

func NewServer(engine *gin.Engine, g *github.Client, port string) *Server {
	return &Server{
		Engine:       engine,
		GithubClient: g,
		Port:         port,
	}
}

func (s *Server) Init() {
	s.Engine.Use(gin.Logger())
	s.Engine.Use(cors.Default())
	s.Engine.POST("/search", s.searchHandler)
}

func (s *Server) Run() error {
	return s.Engine.Run(s.Port)
}

func (s *Server) searchHandler(c *gin.Context) {
	filter := &models.Filter{}
	if err := c.BindJSON(filter); err != nil {
		return
	}
	repos, err := s.GithubClient.ListRepositories()
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"msg": err.Error(),
		})
		return
	}
	data, err := s.GithubClient.GetLanguageAndLines(repos)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(500, gin.H{
			"msg": err.Error(),
		})
		return
	}
	var filteredData []*models.Data
	for _, repo := range data {
		if strings.Contains(repo.Name, filter.Filter) {
			filteredData = append(filteredData, repo)
		}
	}
	c.JSON(http.StatusOK, filteredData)
}
