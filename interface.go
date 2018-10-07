package accuweather

import "time"

type Country struct {
	ID            string `json:"ID"`
	LocalizedName string `json:"LocalizedName"`
	EnglishName   string `json:"EnglishName"`
}

type Region struct {
	ID            string `json:"ID"`
	LocalizedName string `json:"LocalizedName"`
	EnglishName   string `json:"EnglishName"`
}

type AdministrativeArea struct {
	ID            string `json:"ID"`
	LocalizedName string `json:"LocalizedName"`
	EnglishName   string `json:"EnglishName"`
	Level         int    `json:"Level"`
	LocalizedType string `json:"LocalizedType"`
	EnglishType   string `json:"EnglishType"`
	CountryID     string `json:"CountryID"`
}

type TimeZone struct {
	Code             string    `json:"Code"`
	Name             string    `json:"Name"`
	GmtOffset        float32   `json:"GmtOffset"`
	IsDaylightSaving bool      `json:"IsDaylightSaving"`
	NextOffsetChange time.Time `json:"NextOffsetChange"`
}

type Location struct {
	Version            int                `json:"Version"`
	Key                string             `json:"Key"`
	Type               string             `json:"Type"`
	Rank               int                `json:"Rank"`
	LocalizedName      string             `json:"LocalizedName"`
	EnglishName        string             `json:"EnglishName"`
	PrimaryPostalCode  string             `json:"PrimaryPostalCode"`
	Region             Region             `json:"Region"`
	Country            Country            `json:"Country"`
	AdministrativeArea AdministrativeArea `json:"AdministrativeArea"`
	TimeZone           TimeZone           `json:"TimeZone"`
	GeoPosition        struct {
		Latitude  float64 `json:"Latitude"`
		Longitude float64 `json:"Longitude"`
		Elevation struct {
			Metric struct {
				Value    float32 `json:"Value"`
				Unit     string  `json:"Unit"`
				UnitType int     `json:"UnitType"`
			} `json:"Metric"`
			Imperial struct {
				Value    float32 `json:"Value"`
				Unit     string  `json:"Unit"`
				UnitType int     `json:"UnitType"`
			} `json:"Imperial"`
		} `json:"Elevation"`
	} `json:"GeoPosition"`
	IsAlias                bool `json:"IsAlias"`
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
