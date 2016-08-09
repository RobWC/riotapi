package riotapi

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/*
API End Points

BR	BR1	br.api.pvp.net
EUNE	EUN1	eune.api.pvp.net
EUW	EUW1	euw.api.pvp.net
KR	KR	kr.api.pvp.net
LAN	LA1	lan.api.pvp.net
LAS	LA2	las.api.pvp.net
NA	NA1	na.api.pvp.net
OCE	OC1	oce.api.pvp.net
TR	TR1	tr.api.pvp.net
RU	RU	ru.api.pvp.net
PBE	PBE1	pbe.api.pvp.net
*/

// APIEndpoints mapping to api endpoints
var APIEndpoints = map[string]string{
	"br":   "br.api.pvp.net",
	"eune": "eune.api.pvp.net",
	"euw":  "euw.api.pvp.net",
	"kr":   "kr.api.pvp.net",
	"lan":  "lan.api.pvp.net",
	"las":  "las.api.pvp.net",
	"na":   "na.api.pvp.net",
	"oce":  "oce.api.pvp.net",
	"tr":   "tr.api.pvp.net",
	"ru":   "ru.api.pvp.net",
	"jp":   "jp.api.pvp.net",
}

// ShardName a mapping of regions to shard names
var ShardName = map[string]string{
	"br":   "BR1",
	"eune": "EUN1",
	"euw":  "EUW1",
	"kr":   "KR",
	"lan":  "LA1",
	"las":  "LA2",
	"na":   "NA1",
	"oce":  "OC1",
	"tr":   "TR1",
	"ru":   "RU",
	"jp":   "JP1",
}

// RateLimit the current rate limit of the api
type RateLimit struct {
	Limits     map[string]string
	RetryAfter int
}

// APIClient Riot API client
type APIClient struct {
	endpoint      string
	game          string
	region        string
	client        *http.Client
	totalRequests int
	key           string // API key
	tokens        chan struct{}
	RateLimit     *RateLimit
	shardName     string
}

// NewAPIClient create an initalized APIClient
func NewAPIClient(region, key string) *APIClient {
	c := &APIClient{
		key:       key,
		game:      "lol",
		region:    strings.ToLower(region),
		shardName: ShardName[strings.ToLower(region)],
		client: &http.Client{
			Jar:     nil,
			Timeout: time.Second * 5,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				Dial: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 10 * time.Second,
			}},
		tokens:    make(chan struct{}, 20),
		RateLimit: &RateLimit{Limits: make(map[string]string)},
	}

	return c
}

func (c *APIClient) genURL(path []string) string {
	return strings.Join(path, "/")
}

func (c *APIClient) genRequest(method, version, api string, query url.Values) (*http.Request, error) {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = strings.Join([]string{c.region, ".api.pvp.net"}, "")
	u.Path = fmt.Sprintf("/api/%s/%s/%s/%s", c.game, c.region, version, api)
	u.RawQuery = query.Encode()
	return http.NewRequest(method, u.String(), nil)
}

// do execute a request
func (c *APIClient) do(req *http.Request, apiKey bool) ([]byte, error) {
	// add api key

	// manipulate request
	if apiKey {
		query := req.URL.Query()
		query.Add("api_key", strings.ToUpper(c.key))

		//rebuild request
		req.URL.RawQuery = query.Encode()
	}

	if c.RateLimit.RetryAfter > 0 {
		time.Sleep(time.Second * time.Duration(c.RateLimit.RetryAfter))
	}

	rc := make(chan struct {
		data []byte
		err  error
	})

	go func(req http.Request) {
		defer func() {
			<-c.tokens
		}()

		c.tokens <- struct{}{}
		resp, err := c.client.Do(&req)
		if err != nil {
			rc <- struct {
				data []byte
				err  error
			}{nil, err}
		}
		defer resp.Body.Close()

		rl := req.Header.Get("X-Rate-Limit-Count")
		if rl != "" {
			c.updateRateLimits(rl)
		}

		ra := req.Header.Get("Retry-After")
		if rl != "" {
			c.updateRetry(ra)
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			rc <- struct {
				data []byte
				err  error
			}{nil, err}
		}
		if resp.StatusCode != http.StatusOK {
			rc <- struct {
				data []byte
				err  error
			}{nil, fmt.Errorf("API Error %s", string(data))}
		}

		rc <- struct {
			data []byte
			err  error
		}{data, err}
	}(*req)

	r := <-rc
	return r.data, r.err
}
