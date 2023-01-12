package ch7

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

type Track struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Year   int    `json:"year"`
	Length string `json:"length"`
}

type TrackList struct {
	TrackList []*Track `json:"track_list"`
}

func (tL *TrackList) Print() {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

	_, err := fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	if err != nil {
		log.Fatal(err)
	}
	for _, track := range tL.TrackList {
		_, err := fmt.Fprintf(tw, format, track.Title, track.Artist, track.Album, track.Year, track.Length)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tw.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func SortingMusic() {
	var trackList TrackList
	data, err := os.ReadFile("./data/tracks.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &trackList)
	if err != nil {
		log.Fatal(err)
	}

	trackList.Print()

}
