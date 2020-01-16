package service

import "net/http"

//RandomService randome service
type RandomService struct{}

//NewRandomService returns a new random  service
func NewRandomService() *RandomService {
	return &RandomService{}
}

//Get returns response
func (rs *RandomService) Get() (int, error) {
	return http.StatusOK, nil
}
