package common

import (
	"vayu/pkg/ratelimiting_middleware"
	"strings"
	"os"
	"log"
	"vayu/pkg/apperror"
	"net/http"
	"strconv"
)
//RateLimitingHandler handler that ties ratelimiting middleware
type RateLimitingHandler struct{
	ratelimiter ratelimiting_middleware.RateLimiter
}

//NewRateLimitingHandler returns a RateLimitingHandler
func NewRateLimitingHandler() *RateLimitingHandler{
	reqCount:= os.Getenv("REQ_COUNT")
	timeLimit:= os.Getenv("TIME_LIMIT_IN_MINUTES")

	if strings.Trim(reqCount," ")== "" {
		log.Fatal("REQ_COUNT was not provided")
	}

	if strings.Trim(timeLimit," ")=="" {
		log.Fatal("TIME_LIMIT_IN_MINUTES was not provided")
	}
	
	r, countErr:= strconv.ParseFloat(reqCount,64)
	if countErr!=nil {
		log.Fatal("Invalid Request count")
	}

	t, tlimitErr:= strconv.ParseFloat(timeLimit,64)
	if tlimitErr!=nil {
		log.Fatal("Invalid time limit")
	}

	reqPerSec:= getTimePerRequest(r, t)
	log.Println("time to process one request", reqPerSec)

	fullWindowReqCount:= getFullWindowReqCount(r,t)
	return &RateLimitingHandler{
		ratelimiter: ratelimiting_middleware.NewRateLimitingProvider(reqPerSec, fullWindowReqCount),
	}
}

//RateLimitHandle handler for ratelimiter
func(rh *RateLimitingHandler) RateLimitHandle(next http.Handler) http.Handler{
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request){
			if !rh.ratelimiter.RateLimit(){
				SendErrorResponse(w,r,apperror.NewAppError("# of requests more than limit", http.StatusTooManyRequests))
				return
			}

			next.ServeHTTP(w,r)
		})
}


//getTimePerRequest - # of seconds for one request in an evenly distributed window
func getTimePerRequest(r, t float64) float64 {
	secondsInMinute:= float64(60)
	return t * secondsInMinute/ r
}


func getFullWindowReqCount(r,t float64) float64{
	secondsInMinute:= float64(60)
	return secondsInMinute * t
}