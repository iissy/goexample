package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/olivere/elastic/v6"
	"github.com/sirupsen/logrus"
)

var client *ElasticClient

func init() {
	var err error
	client, err = NewElasticClient(ElasticClientConfig{
		Addr:     "http://192.168.1.234:9200",
		User:     "",
		Password: "",
	})
	if err != nil {
		log.Panic(err)
	}
}

type ReqEsMatch struct {
	Match map[string]string `json:"match"`
}
type ReqEsTerm struct {
	Terms map[string][]string `json:"terms"`
}
type ReqEsBool struct {
	Must []interface{} `json:"must"`
}

func EsQueryFrame(warehouseID string, containerIdList []string) (frameList []EsResFrameSourceDetail, err error) {
	//查询条件
	warehosueIDMatch := ReqEsMatch{}
	warehosueIDMatch.Match = map[string]string{"warehouse_id": warehouseID}
	termContainerId := ReqEsTerm{}
	termContainerId.Terms = map[string][]string{"frame_id": containerIdList}

	queryBool := ReqEsBool{}
	queryBool.Must = []interface{}{warehosueIDMatch, termContainerId}

	//es查询
	frameIndex := "wareservice.frame"
	scrollID := ""
	for {
		res, err := client.Scroll(frameIndex, ElasticScrollRequest{
			ScrollID: scrollID,
			ElasticSearchRequest: ElasticSearchRequest{
				Query: map[string]interface{}{
					"query": map[string]interface{}{"bool": queryBool},
					"size":  10000,
				},
			},
		})
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic(err)
		}
		scrollID = res.ScrollId
		esResHits := res.Hits.Hits
		for _, esResSource := range esResHits {
			var frame EsResFrameSourceDetail
			err := json.Unmarshal(*esResSource.Source, &frame)
			if err != nil {
				log.Panic(err)
			}
			frameList = append(frameList, frame)
		}
	}

	return frameList, nil
}

type EsResFrameSourceDetail struct {
	FrameID      string `json:"frame_id"`
	Node         string `json:"node"`
	PositionType int    `json:"position_type"`
	CurrentZ     int    `json:"current_z"`
}

func main() {
	for i := 1; i < 10000; i++ {
		list, err := EsQueryFrame("232414789256686833", []string{"354956860203204610"})
		if err != nil {
			log.Panic(err)
		}

		for _, item := range list {
			log.Printf("%+v", item)
		}
	}
}

type ElasticClient struct {
	c *elastic.Client
}

func (c *ElasticClient) GetRawClient() *elastic.Client {
	return c.c
}

type ElasticScrollRequest struct {
	ElasticSearchRequest
	KeepaliveTime string
	ScrollID      string
}

func (c *ElasticClient) Scroll(index string, req ElasticScrollRequest) (*elastic.SearchResult, error) {
	doCtx := context.Background()
	keepaliveTime := "1m"
	if req.KeepaliveTime != "" {
		keepaliveTime = req.KeepaliveTime
	}
	scrollService := c.c.Scroll().KeepAlive(keepaliveTime).Index(index)
	if req.ScrollID != "" {
		scrollService = scrollService.ScrollId(req.ScrollID)
	}
	searchResult, err := scrollService.Body(req.Query).Do(doCtx)
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}

type ElasticSearchRequest struct {
	Query interface{}
}

func (c *ElasticClient) Search(index string, req ElasticSearchRequest) (*elastic.SearchResult, error) {
	doCtx := context.Background()
	searchService := c.c.Search().Index(index)
	searchResult, err := searchService.Source(req.Query).Do(doCtx)
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}

// todo use sql in es
func (c *ElasticClient) Sql(index string) (*elastic.SearchResult, error) {
	return nil, errors.New("not ready yet")
}

type ElasticClientConfig struct {
	HTTPS         bool
	Addr          string
	User          string
	Password      string
	Decoder       elastic.Decoder
	RetryStrategy elastic.Retrier
	HttpClient    *http.Client
	Other         []elastic.ClientOptionFunc
	LogLevel      string
}

func NewElasticClient(cfg ElasticClientConfig) (*ElasticClient, error) {
	var options []elastic.ClientOptionFunc
	options = append(options, elastic.SetURL(cfg.Addr))
	options = append(options, elastic.SetSniff(false))
	options = append(options, elastic.SetGzip(true))
	if cfg.User != "" || cfg.Password != "" {
		options = append(options, elastic.SetBasicAuth(cfg.User, cfg.Password))
	}
	if cfg.HTTPS {
		options = append(options, elastic.SetScheme("https"))
	}
	if cfg.HttpClient != nil {
		options = append(options, elastic.SetHttpClient(cfg.HttpClient))
	}
	if cfg.Decoder != nil {
		options = append(options, elastic.SetDecoder(cfg.Decoder))
	}

	if cfg.RetryStrategy != nil {
		options = append(options, elastic.SetRetrier(cfg.RetryStrategy))
	}

	loglevel := logrus.InfoLevel
	if level, exist := getLogLevel(cfg.LogLevel); exist {
		loglevel = level
	}
	elasticLogger := newElasticLogger(loglevel)
	options = append(options, elastic.SetErrorLog(&elasticErrorLogger{elasticLogger}))
	options = append(options, elastic.SetTraceLog(&elasticDebugLogger{elasticLogger}))
	options = append(options, elastic.SetInfoLog(&elasticInfoLogger{elasticLogger}))

	// override
	if len(cfg.Other) > 0 {
		options = append(options, cfg.Other...)
	}
	es, err := elastic.NewClient(options...)
	if err != nil {
		return nil, err
	}

	return &ElasticClient{
		c: es,
	}, nil
}

type elasticLogger struct {
	logger *logrus.Logger
}

func newElasticLogger(level logrus.Level) elasticLogger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return elasticLogger{
		logger: logger,
	}
}

func (l *elasticLogger) getLogger() *logrus.Entry {
	return l.logger.WithFields(logrus.Fields{"module": "elastic"})
}

type elasticDebugLogger struct {
	elasticLogger
}
type elasticInfoLogger struct {
	elasticLogger
}
type elasticErrorLogger struct {
	elasticLogger
}

func (l *elasticDebugLogger) Printf(format string, v ...interface{}) {
	l.getLogger().Debugf(format, v...)
}

func (l *elasticInfoLogger) Printf(format string, v ...interface{}) {
	l.getLogger().Infof(format, v...)
}

func (l *elasticErrorLogger) Printf(format string, v ...interface{}) {
	l.getLogger().Errorf(format, v...)
}

func getLogLevel(loglevel string) (logrus.Level, bool) {
	switch loglevel {
	case "debug":
		return logrus.DebugLevel, true
	case "info":
		return logrus.InfoLevel, true
	case "warn":
		return logrus.WarnLevel, true
	case "error":
		return logrus.ErrorLevel, true
	case "fatal":
		return logrus.FatalLevel, true
	case "panic":
		return logrus.PanicLevel, true
	default:
		return 0, false
	}
}
