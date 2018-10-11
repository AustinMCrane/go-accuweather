package accuweather

// Country country of origin
type Country struct {
	ID            string `json:"id"`
	LocalizedName string `json:"localizedName"`
	EnglishName   string `json:"englishName"`
}

// Region region of origin
type Region struct {
	ID            string `json:"id"`
	LocalizedName string `json:"localizedName"`
	EnglishName   string `json:"englishName"`
}

// AdministrativeArea area of origin
type AdministrativeArea struct {
	ID            string `json:"id"`
	LocalizedName string `json:"localizedName"`
	EnglishName   string `json:"englishName"`
	Level         int    `json:"level"`
	LocalizedType string `json:"localizedType"`
	EnglishType   string `json:"englishType"`
	CountryID     string `json:"countryID"`
}

// GeoPosition geographical position
type GeoPosition struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Elevation struct {
		Metric   Measure `json:"metric"`
		Imperial Measure `json:"imperial"`
	} `json:"elevation"`
}

// Location is a key part of the api, you must have a location
// to get any of the data with a location key
type Location struct {
	Version int `json:"version"`
	// Key is the key to the location for anything that is specific to a location
	Key                    string             `json:"key"`
	Type                   string             `json:"type"`
	Rank                   int                `json:"rank"`
	LocalizedName          string             `json:"localizedName"`
	EnglishName            string             `json:"englishName"`
	PrimaryPostalCode      string             `json:"primaryPostalCode"`
	Region                 Region             `json:"region"`
	Country                Country            `json:"country"`
	AdministrativeArea     AdministrativeArea `json:"administrativeArea"`
	TimeZone               TimeZone           `json:"timeZone"`
	GeoPosition            GeoPosition        `json:"GeoPosition"`
	IsAlias                bool               `json:"IsAlias"`
	SupplementalAdminAreas []struct {
		Level         int    `json:"Level"`
		LocalizedName string `json:"LocalizedName"`
		EnglishName   string `json:"EnglishName"`
	} `json:"SupplementalAdminAreas"`
	DataSets []string `json:"DataSets"`
	Details  struct {
		Key                      string      `json:"Key"`
		StationCode              string      `json:"StationCode"`
		StationGmtOffset         float32     `json:"StationGmtOffset"`
		BandMap                  string      `json:"BandMap"`
		Climo                    string      `json:"Climo"`
		LocalRadar               string      `json:"LocalRadar"`
		MediaRegion              string      `json:"MediaRegion"`
		Metar                    string      `json:"Metar"`
		NXMetro                  string      `json:"NXMetro"`
		NXState                  string      `json:"NXState"`
		Population               int         `json:"Population"`
		PrimaryWarningCountyCode string      `json:"PrimaryWarningCountyCode"`
		PrimaryWarningZoneCode   string      `json:"PrimaryWarningZoneCode"`
		Satellite                string      `json:"Satellite"`
		Synoptic                 string      `json:"Synoptic"`
		MarineStation            string      `json:"MarineStation"`
		MarineStationGMTOffset   interface{} `json:"MarineStationGMTOffset"`
		VideoCode                string      `json:"VideoCode"`
		LocationStem             string      `json:"LocationStem"`
		DMA                      struct {
			ID          string `json:"ID"`
			EnglishName string `json:"EnglishName"`
		} `json:"DMA"`
		PartnerID interface{} `json:"PartnerID"`
		Sources   []struct {
			DataType string `json:"DataType"`
			Source   string `json:"Source"`
			SourceID int    `json:"SourceId"`
		} `json:"Sources"`
		CanonicalPostalCode  string `json:"CanonicalPostalCode"`
		CanonicalLocationKey string `json:"CanonicalLocationKey"`
	} `json:"Details"`
}
