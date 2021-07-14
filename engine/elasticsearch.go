package engine

import (
  "fmt"
  "github.com/olivere/elastic/v7"
  "log"
  "os"
  "shopping/config"
)

var esClient *elastic.Client

func init()  {
  esConfig := config.GetElasticSearchConfig()
  url := fmt.Sprintf("http://%s:%s", esConfig.Host, esConfig.Port)

  client, err := elastic.NewClient(
    elastic.SetURL(url),
    elastic.SetSniff(false),
    elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC", log.LstdFlags)),
    elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
  if err != nil {
    panic(fmt.Errorf("Connet elasticsearch failed, err: %s \n", err))
  }

  esClient = client
}

func GetEsClient() *elastic.Client  {
  return esClient
}
