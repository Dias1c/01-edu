package api

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

// Refresh data in AllArtists
func RefreshAll() {
	if !GetNewData() {
		log.Print("Can't refresh Artists cause something wrong with API")
		return
	}
	rewriteCache(NewAllArtists)
	*AllArtists = make([]Artists, len(*NewAllArtists))
	copy(*AllArtists, *NewAllArtists)
}

// Get new data from API
func GetNewData() bool {
	// Flush NewAllArtists
	NewAllArtists = &[]Artists{}
	// At first parse artists data
	var artists []artistsJSON
	var relations relationsJSON
	success := parserAPI(artistsAPI, &artists)
	if !success {
		return success
	}
	success = parserAPI(relationAPI, &relations)
	if !success {
		return success
	}
	// For every artist get relation info and append it to All
	for ind, artist := range artists {
		tempArtist := Artists{
			Id:                  artist.Id,
			Name:                artist.Name,
			Image:               artist.Image,
			Members:             artist.Members,
			CreationDate:        artist.CreationDate,
			FirstAlbumFormatted: artist.FirstAlbum,
			Concerts:            []Concert{},
		}
		date, err := time.Parse(timeLayout, artist.FirstAlbum)
		if err != nil {
			log.Print(err.Error())
		}
		tempArtist.FirstAlbum = date
		// Parse relation
		// Maybe change after (to do logic)
		// if artist.Id == relations.Index[ind].Id {
		parserConcert(&tempArtist.Concerts, relations.Index[ind].DatesLocations)
		// }
		*NewAllArtists = append(*NewAllArtists, tempArtist)
	}
	return true
}

// Get suggestions for autofill or suggestion
func GetSuggestions() {
	AllSuggestions = &[]Suggestions{}
	temp := make(map[Suggestions]bool)
	for _, artist := range *AllArtists {
		temp[Suggestions{Name: artist.Name, Type: "Name"}] = true
		for _, member := range artist.Members {
			temp[Suggestions{Name: member, Type: "Member"}] = true
		}
		temp[Suggestions{Name: strconv.Itoa(artist.CreationDate), Type: "Creation Date"}] = true
		temp[Suggestions{Name: artist.FirstAlbumFormatted, Type: "First Album Date"}] = true
		for _, concerts := range artist.Concerts {
			temp[Suggestions{Name: concerts.Location, Type: "Concert Location"}] = true
		}
	}
	for key := range temp {
		*AllSuggestions = append(*AllSuggestions, key)
	}
	sort.Slice(*AllSuggestions, func(i, j int) bool {
		return (*AllSuggestions)[i].Type < (*AllSuggestions)[j].Type
	})
}

// MAP API : Gives formatted location name and coordinates by the city name
func geocoding(city string, concert *Concert) bool {
	url := fmt.Sprintf(geocodingURL, city, googleMapAPI)
	var data geocodingJSON
	success := parserAPI(url, &data)
	if !success || len(data.Results) < 1 {
		return false
	}
	concert.Location = data.Results[0].Address
	concert.Coordinates[0] = data.Results[0].Geometry.Location.Latitude
	concert.Coordinates[1] = data.Results[0].Geometry.Location.Longtitude
	return true
}
