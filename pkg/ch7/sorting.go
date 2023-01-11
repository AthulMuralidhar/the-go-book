package ch7

import "time"

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func (t Track) UnmarshalJSON(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}

func (t Track) String() string {
	//TODO implement me
	panic("implement me")
}

type byArtist []*Track

func (tList byArtist) Len() int {
	return len(tList)
}

func (tList byArtist) Less(i, j int) bool {
	return tList[i].Artist < tList[j].Artist
}

func (tList byArtist) Swap(i, j int) {
	tList[i], tList[j] = tList[j], tList[i]
}
