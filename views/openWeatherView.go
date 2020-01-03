package views

// OpenWeather ...
type OpenWeather struct {
	Coord struct {
		Lon float32 `json:"lon,omitempty"`
		Lat float32 `json:"lat,omitempty"`
	} `coord:"coord,omitempty"`
	Weather []struct {
		ID          int    `json:"id,omitempty"`
		Main        string `json:"main,omitempty"`
		Description string `json:"description,omitempty"`
		Icon        string `json:"icon,omitempty"`
	} `json:"weather,omitempty"`
	Base string `json:"base,omitempty"`
	Main struct {
		Temp     float32 `json:"temp,omitempty"`
		Pressure int     `json:"pressure,omitempty"`
		Humidity int     `json:"humidity,omitempty"`
		TempMin  float32 `json:"temp_min,omitempty"`
		TempMax  float32 `json:"temp_max,omitempty"`
	} `json:"main,omitempty"`
	Visibility int `json:"visibility,omitempty"`
	Wind       struct {
		Speed float32 `json:"speed,omitempty"`
		Deg   int     `json:"deg,omitempty"`
	} `json:"wind,omitempty"`
	Clouds struct {
		All int `json:"all,omitempty"`
	} `json:"clouds,omitempty"`
	Dt  int `json:"dt,omitempty"`
	Sys struct {
		Type    int     `json:"type,omitempty"`
		ID      int     `json:"id,omitempty"`
		Message float32 `json:"message,omitempty"`
		Country string  `json:"country,omitempty"`
		Sunrise int     `json:"sunrise,omitempty"`
		Sunset  int     `json:"sunset,omitempty"`
	} `json:"sys,omitempty"`
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Cod  int    `json:"cod,omitempty"`
}

// OpenWeatherForecast ...
type OpenWeatherForecast struct {
	Cod     string        `json:"cod,omitempty"`
	Message int           `json:"message,omitempty"`
	Cnt     int           `json:"cnt,omitempty"`
	List    []OpenWeather `json:"list,omitempty"`
	City    struct {
		ID    int    `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Coord struct {
			Lon float32 `json:"lon,omitempty"`
			Lat float32 `json:"lat,omitempty"`
		} `coord:"coord,omitempty"`
		Country    string `json:"country,omitempty"`
		Population int    `json:"population,omitempty"`
		Timezone   int    `json:"timezone,omitempty"`
		Sunrise    int    `json:"sunrise,omitempty"`
		Sunset     int    `json:"sunset,omitempty"`
	}
}

// QueryResult ...
type QueryResult struct {
	QueryResult struct {
		Action     string `json:"action,omitempty"`
		Parameters struct {
			City      string `json:"City,omitempty"`
			When      string `json:"When,omitempty"`
			Firstname string `json:"Firstname,omitempty"`
			Subject   string `json:"Subject,omitempty"`
		} `json:"Parameters"`
	} `json:"queryResult"`
}

// FollowupCityEvent ...
type FollowupCityEvent struct {
	FollowupEventInput FollowupEventInput `json:"followupEventInput,omitempty"`
}

// FollowupEventInput ...
type FollowupEventInput struct {
	Name string `json:"name,omitempty"`
}
