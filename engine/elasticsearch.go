package engine

import (
  "context"
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
    elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC", log.LstdFlags)),
    elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
  if err != nil {
    panic(fmt.Errorf("Connet elasticsearch failed, err: %s \n", err))
  }

  esClient = client
}

func GetEsClient() *elastic.Client  {
  esConfig := config.GetElasticSearchConfig()
  url := fmt.Sprintf("http://%s:%s", esConfig.Host, esConfig.Port)
  fmt.Printf("ES的url是: %s\n", url)
  info, code, err := esClient.Ping(url).Do(context.Background())
  if err != nil {
    panic(err)
  }
  fmt.Printf("ES returned with code: %d, version: %s", code, info.Version.Number)

  return esClient
}
