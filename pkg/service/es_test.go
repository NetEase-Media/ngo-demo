package service

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/NetEase-Media/ngo/adapter/log"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/stretchr/testify/assert"
)

var ESClient = InitElasticsearch()

func InitElasticsearch() *elasticsearch.Client {
	//elasticsearchUrl := common.GetStringFromConfig("elasticsearch")
	//if elasticsearchUrl == "" {
	//	return nil
	//}
	elasticsearchUrl := ""
	cfg := elasticsearch.Config{
		Addresses: []string{
			elasticsearchUrl,
		}, Header: http.Header{
			"index": []string{"elephant_playlist_open"},
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.WithField("err", err).Error("elasticsearch初始化失败")
	}
	return es
}
func TestESinfo(t *testing.T) {

}

func TestSearch(t *testing.T) {
	a := assert.New(t)
	body := &bytes.Buffer{}
	body.WriteString(`
	{
		"_source":{
		  "excludes": ["name"]
		}, 
		"query": {
		  "match_phrase": {
			"name": "沙特"
		  }
		}, 
		"from": 0,
		"size": 5
	}
	`)
	response, err := ESClient.Search(ESClient.Search.WithIndex("elephant_playlist_open"), ESClient.Search.WithBody(body))
	a.Nil(err)
	t.Log(response)
}
