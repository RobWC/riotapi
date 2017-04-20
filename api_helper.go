package riotapi

import (
	"log"
	"strconv"
	"strings"
	"time"
)

// updateRateLimits update the current rate limits
func (c *APIClient) updateRateLimits(limits string) {

	if len(limits) > 0 {
		values := strings.Split(limits, ",")

		for v := range values {
			nv := strings.Split(values[v], ":")
			if len(nv) == 2 {
				c.RateLimit.Limits[nv[0]] = nv[1]
			}
		}
	}
}

// updateRetry update the current retry wait time
func (c *APIClient) updateRetry(retry string) {
	v, err := strconv.Atoi(retry)
	if err != nil {
		// retry value incorrect
		log.Printf("Retry value incorrect %s\n", retry)
	}
	c.RateLimit.RetryAfter = v
}

// updateRateLimitCount
func (c *APIClient) updateRateLimitCount(count string) {

}
