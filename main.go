package main

import (
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"encoding/json" // to convert the response body to struct

	"io" // to read the response body

	"github.com/gin-gonic/gin"
)

// Enum
type Implementation int

const (
	CSV Implementation = iota
	API
	MOCK
)

type Service interface {
	Distance(city1 string, city2 string) (int, error)
}

type (
	CSVService  struct{}
	APIService  struct{}
	MockService struct{}
)

func buildUrl(args ...string) string {
	var urlBuilder strings.Builder
	for _, arg := range args {
		urlBuilder.WriteString(arg)
	}
	return urlBuilder.String()
}

func (s *CSVService) Distance(city1 string, city2 string) (int, error) {
	return 1, nil
}

type APIResponse struct {
	Name string `json:"name"`
}

type Place struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func (s *APIService) Distance(city1 string, city2 string) (int, error) {
	// https://nominatim.openstreetmap.org/search?q=lima,peru&format=json

	url1 := buildUrl("https://nominatim.openstreetmap.org/search?q=", city1, "&format=json")
	url2 := buildUrl("https://nominatim.openstreetmap.org/search?q=", city2, "&format=json")

	var err1 error
	var err2 error

	response1, err1 := http.Get(url1)
	response2, err2 := http.Get(url2)

	if err1 != nil {
		log.Println(err1)
		return 0, err1
	}
	if err2 != nil {
		log.Println(err2)
		return 0, err2
	}

	defer response1.Body.Close()
	defer response2.Body.Close()

	body1, err1 := io.ReadAll(response1.Body)
	body2, err2 := io.ReadAll(response2.Body)

	if err1 != nil {
		log.Println(err1)
		return 0, err1
	}
	if err2 != nil {
		log.Println(err2)
		return 0, err2
	}

	places1 := make([]Place, 0)
	err1 = json.Unmarshal(body1, &places1)

	places2 := make([]Place, 0)
	err2 = json.Unmarshal(body2, &places2)

	if err1 != nil {
		log.Println(err1)
		return 0, err1
	}
	if err2 != nil {
		log.Println(err2)
		return 0, err2
	}

	return computeDistance(places1[0].Lat, places1[0].Lon, places2[0].Lat, places2[0].Lon), nil
}

func computeDistance(lat1 string, lon1 string, lat2 string, lon2 string) int {
	earth_radius := 6371.0

	rad_lat1, _ := strconv.ParseFloat(lat1, 64)
	rad_lon1, _ := strconv.ParseFloat(lon1, 64)
	rad_lat2, _ := strconv.ParseFloat(lat2, 64)
	rad_lon2, _ := strconv.ParseFloat(lon2, 64)

	rad_lat1 = rad_lat1 * math.Pi / 180
	rad_lon1 = rad_lon1 * math.Pi / 180
	rad_lat2 = rad_lat2 * math.Pi / 180
	rad_lon2 = rad_lon2 * math.Pi / 180

	diff_lat := rad_lat2 - rad_lat1
	diff_lon := rad_lon2 - rad_lon1

	a := math.Sin(diff_lat/2)*math.Sin(diff_lat/2) + math.Cos(rad_lat1)*math.Cos(rad_lat2)*math.Sin(diff_lon/2)*math.Sin(diff_lon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return int(earth_radius * c)
}

func (s *MockService) Distance(city1 string, city2 string) (int, error) {
	return rand.Intn(100), nil
}

func NewService(method string) Service {
	switch method {
	case "CSV":
		return &CSVService{}
	case "API":
		return &APIService{}
	case "MOCK":
		return &MockService{}
	default:
		return nil
	}
}

func ginService(gin_ctx *gin.Context) {
	gin_ctx.Header("Access-Control-Allow-Origin", "*")
	gin_ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	city1 := gin_ctx.Query("city1")
	city2 := gin_ctx.Query("city2")

	method := gin_ctx.Query("method")

	log.Println("method: ", method)

	// Create a new service
	service := NewService(method)

	// Get the distance
	distance, error := service.Distance(city1, city2)

	if error != nil {
		log.Println("error: ", error)
		gin_ctx.JSON(500, gin.H{"error": error})
	}

	log.Println("distance: ", distance)

	gin_ctx.JSON(200, gin.H{"distance": distance})
}

func main() {
	router := gin.Default()

	router.POST("/distance", ginService)

	router.Run(":3003")
}
