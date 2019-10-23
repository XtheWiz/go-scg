package main

import (
	"math"
	"time"
)

type ReturnPlaceResult struct {
	Name     string
	PhotoRef string
	Lat      float64
	Lng      float64
	Distance float64
	Vicinity string
}

type PlacesSearchResponse struct {
	// Results is the Place results for the search query
	Results []PlacesSearchResult
	// HTMLAttributions contain a set of attributions about this listing which must be
	// displayed to the user.
	HTMLAttributions []string
	// NextPageToken contains a token that can be used to return up to 20 additional
	// results.
	NextPageToken string

	Status string
}

type PlacesSearchResult struct {
	// FormattedAddress is the human-readable address of this place
	FormattedAddress string `json:"formatted_address,omitempty"`
	// Geometry contains geometry information about the result, generally including the
	// location (geocode) of the place and (optionally) the viewport identifying its
	// general area of coverage.
	Geometry AddressGeometry `json:"geometry,omitempty"`
	// Name contains the human-readable name for the returned result. For establishment
	// results, this is usually the business name.
	Name string `json:"name,omitempty"`
	// Icon contains the URL of a recommended icon which may be displayed to the user
	// when indicating this result.
	Icon string `json:"icon,omitempty"`
	// PlaceID is a textual identifier that uniquely identifies a place.
	PlaceID string `json:"place_id,omitempty"`
	// Scope indicates the scope of the PlaceID.
	Scope string `json:"scope,omitempty"`
	// Rating contains the place's rating, from 1.0 to 5.0, based on aggregated user
	// reviews.
	Rating float32 `json:"rating,omitempty"`
	// UserRatingsTotal contains total number of the place's ratings
	UserRatingsTotal int `json:"user_ratings_total,omitempty"`
	// Types contains an array of feature types describing the given result.
	Types []string `json:"types,omitempty"`
	// OpeningHours may contain whether the place is open now or not.
	OpeningHours *OpeningHours `json:"opening_hours,omitempty"`
	// Photos is an array of photo objects, each containing a reference to an image.
	Photos []Photo `json:"photos,omitempty"`
	// AltIDs — An array of zero, one or more alternative place IDs for the place, with
	// a scope related to each alternative ID.
	AltIDs []AltID `json:"alt_ids,omitempty"`
	// PriceLevel is the price level of the place, on a scale of 0 to 4.
	PriceLevel int `json:"price_level,omitempty"`
	// Vicinity contains a feature name of a nearby location.
	Vicinity string `json:"vicinity,omitempty"`
	// PermanentlyClosed is a boolean flag indicating whether the place has permanently
	// shut down.
	PermanentlyClosed bool `json:"permanently_closed,omitempty"`
	// ID is an identifier.
	ID string `json:"id,omitempty"`
}

type AddressGeometry struct {
	Location     LatLng       `json:"location"`
	LocationType string       `json:"location_type"`
	Bounds       LatLngBounds `json:"bounds"`
	Viewport     LatLngBounds `json:"viewport"`
	Types        []string     `json:"types"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type LatLngBounds struct {
	NorthEast LatLng `json:"northeast"`
	SouthWest LatLng `json:"southwest"`
}

type OpeningHours struct {
	// OpenNow is a boolean value indicating if the place is open at the current time.
	// Please note, this field will be null if it isn't present in the response.
	OpenNow *bool `json:"open_now,omitempty"`
	// Periods is an array of opening periods covering seven days, starting from Sunday,
	// in chronological order.
	Periods []OpeningHoursPeriod `json:"periods,omitempty"`
	// weekdayText is an array of seven strings representing the formatted opening hours
	// for each day of the week, for example "Monday: 8:30 am – 5:30 pm".
	WeekdayText []string `json:"weekday_text,omitempty"`
	// PermanentlyClosed indicates that the place has permanently shut down. Please
	// note, this field will be null if it isn't present in the response.
	PermanentlyClosed *bool `json:"permanently_closed,omitempty"`
}

type OpeningHoursPeriod struct {
	// Open is when the place opens.
	Open OpeningHoursOpenClose `json:"open"`
	// Close is when the place closes.
	Close OpeningHoursOpenClose `json:"close"`
}

type OpeningHoursOpenClose struct {
	// Day is a number from 0–6, corresponding to the days of the week, starting on
	// Sunday. For example, 2 means Tuesday.
	Day time.Weekday `json:"day"`
	// Time contains a time of day in 24-hour hhmm format. Values are in the range
	// 0000–2359. The time will be reported in the place’s time zone.
	Time string `json:"time"`
}

type Photo struct {
	// PhotoReference is used to identify the photo when you perform a Photo request.
	PhotoReference string `json:"photo_reference"`
	// Height is the maximum height of the image.
	Height int `json:"height"`
	// Width is the maximum width of the image.
	Width int `json:"width"`
	// htmlAttributions contains any required attributions.
	HTMLAttributions []string `json:"html_attributions"`
}

type AltID struct {
	// PlaceID is the APP scoped Place ID that you received when you initially created
	// this Place, before it was given a Google wide Place ID.
	PlaceID string `json:"place_id,omitempty"`
	// Scope is the scope of this alternative place ID. It will always be APP,
	// indicating that the alternative place ID is recognised by your application only.
	Scope string `json:"scope,omitempty"`
}

//ref: https://socketloop.com/tutorials/golang-how-to-calculate-the-distance-between-two-coordinates-using-haversine-formula
const radius = 6371 // Earth's mean radius in kilometers

func degrees2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

//findDistance : Assume we are at
//Lat = 13.8035134
//Lng = 100.5373821
func (dest ReturnPlaceResult) findDistance() float64 {
	degreesLat := degrees2radians(dest.Lat - 13.8035134)
	degreesLong := degrees2radians(dest.Lng - 100.5373821)
	a := (math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
		math.Cos(degrees2radians(13.8035134))*
			math.Cos(degrees2radians(dest.Lat))*math.Sin(degreesLong/2)*
			math.Sin(degreesLong/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := radius * c

	return math.Round(d*100) / 100
}
