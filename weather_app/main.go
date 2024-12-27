// main.go
package main

import (
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type WeatherData struct {
    Location    string  `json:"location"`
    Temperature float64 `json:"temperature"`
    Description string  `json:"description"`
    Humidity    int     `json:"humidity"`
    WindSpeed   float64 `json:"wind_speed"`
    Icon        string  `json:"icon"`
}

type LocationSuggestion struct {
    Name      string  `json:"name"`
    Latitude  float64 `json:"lat"`
    Longitude float64 `json:"lon"`
}

type WeatherResponse struct {
    Name string `json:"name"`
    Main struct {
        Temp     float64 `json:"temp"`
        Humidity int     `json:"humidity"`
    } `json:"main"`
    Weather []struct {
        Description string `json:"description"`
        Icon        string `json:"icon"`
    } `json:"weather"`
    Wind struct {
        Speed float64 `json:"speed"`
    } `json:"wind"`
}

func main() {
    r := mux.NewRouter()

    // Serve static files
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    
    // Routes
    r.HandleFunc("/", homeHandler).Methods("GET")
    r.HandleFunc("/weather/{location}", weatherHandler).Methods("GET")
    r.HandleFunc("/suggest/{query}", suggestHandler).Methods("GET")

    fmt.Println("Server starting on :8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, nil)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    location := vars["location"]
    
    // Replace with your actual API key
    apiKey := "32230e101350521281fd66409ce7eaaf"
    
    // Make API call to OpenWeatherMap
    url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", location, apiKey)
    resp, err := http.Get(url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    var weatherResp WeatherResponse
    if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Transform the response into our WeatherData struct
    weatherData := WeatherData{
        Location:    weatherResp.Name,
        Temperature: weatherResp.Main.Temp,
        Humidity:    weatherResp.Main.Humidity,
        WindSpeed:   weatherResp.Wind.Speed,
    }

    if len(weatherResp.Weather) > 0 {
        weatherData.Description = weatherResp.Weather[0].Description
        weatherData.Icon = weatherResp.Weather[0].Icon
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(weatherData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func suggestHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    query := vars["query"]
    
    // Replace with your actual API key
    apiKey := "32230e101350521281fd66409ce7eaaf"
    
    url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s", query, apiKey)
    resp, err := http.Get(url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    var suggestions []LocationSuggestion
    if err := json.NewDecoder(resp.Body).Decode(&suggestions); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(suggestions)
}
