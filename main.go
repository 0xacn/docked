package main

import (
  "context"
  "encoding/json"
  "fmt"
  "net/http"
  "net/url"
  "os"
  "os/exec"
  "strings"
  "regexp"

  "docker.io/go-docker"
  "docker.io/go-docker/api/types/registry"
  "dockerio/go-docker/api/types"
)

var dockerClient *docker.Client
var shellCommands commands.Commands = commands.New()

// docker results (api req)
type DockerHubResulttruct {
  ResultCount *int `json:"num_results,omitempty"`
  PageCount *int `json:"num_pages,omitempty"`
  ItemCountPerPage *int `json:"page_size,omitempty"`
  Query *string `json:"query,omitempty"`
  CurrentPage *int `json:"page,omitempty"`
  Items []registry.SearchResult `json:"results,omitempty"`
}
