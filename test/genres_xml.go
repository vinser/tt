package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func (fbgt *FBGenresTransfer) String() string {
	s := ""
	for _, g := range fbgt.Genres {
		s += fmt.Sprintln("genre:", g.Value)
		for _, rd := range g.Descriptions {
			s += fmt.Sprintf("\troot-descr: lang=%s title=%s detailed=%s\n", rd.Lang, rd.Title, rd.Detailed)
		}
		s += fmt.Sprintln("\tsubgenres ---")
		for _, sg := range g.Subgenres {
			s += fmt.Sprintln("\t\tsubgenre: value=", sg.Value)
			for _, sgd := range sg.Descriptions {
				s += fmt.Sprintf("\t\t\tgenre-descr lang=%s title=%s\n", sgd.Lang, sgd.Title)
			}
			for _, sga := range sg.Alts {
				s += fmt.Sprintf("\t\t\tgenre-alt value=%s format=%s\n", sga.Value, sga.Format)
			}
		}
	}
	return s
}

type FBGenresTransfer struct {
	XMLName xml.Name `xml:"fbgenrestransfer"`
	Genres  []Genre  `xml:"genre"`
}

func (fbgt *FBGenresTransfer) Transfer(genre string) string {
	for _, g := range fbgt.Genres {
		for _, sg := range g.Subgenres {
			if sg.Value == genre {
				return sg.Value
			}
			for _, sga := range sg.Alts {
				if sga.Value == genre {
					return sg.Value
				}
			}
		}
	}
	return "no_name:" + genre
}

func (fbgt *FBGenresTransfer) ListGenres() []Genre {
	return fbgt.Genres
}

func (fbgt *FBGenresTransfer) ListSubGenres(genre string) []Subgenre {
	sg := []Subgenre{}
	for _, g := range fbgt.Genres {
		if g.Value == genre {
			return g.Subgenres
		}
	}
	return sg
}

func (fbgt *FBGenresTransfer) GenreName(genre, lang string) string {
	for _, g := range fbgt.Genres {
		for _, sg := range g.Subgenres {
			if sg.Value == genre {
				for _, sgd := range sg.Descriptions {
					if sgd.Lang == lang {
						return sgd.Title
					}
				}
			}
		}
	}
	return ""
}

type Genre struct {
	Value        string      `xml:"value,attr"`
	Descriptions []RootDescr `xml:"root-descr"`
	Subgenres    []Subgenre  `xml:"subgenres>subgenre"`
}

type RootDescr struct {
	Lang     string `xml:"lang,attr"`
	Title    string `xml:"genre-title,attr"`
	Detailed string `xml:"detailed,attr"`
}

type Subgenre struct {
	Value        string       `xml:"value,attr"`
	Descriptions []GenreDescr `xml:"genre-descr"`
	Alts         []GenreAlt   `xml:"genre-alt"`
}

type GenreDescr struct {
	Lang  string `xml:"lang,attr"`
	Title string `xml:"title,attr"`
}

type GenreAlt struct {
	Value  string `xml:"value,attr"`
	Format string `xml:"format,attr"`
}

func main() {
	var err error
	xmlStream, err := os.Open("genres_transfer.xml")
	if err != nil {
		log.Println("Failed to open XML file")
		return
	}
	defer xmlStream.Close()

	xmlData, err := ioutil.ReadAll(xmlStream)

	fbgt := &FBGenresTransfer{}

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.Strict = false
	decoder.Decode(&fbgt)
	// fmt.Println(fbgt)
	// fmt.Println(fbgt.ListSubGenres("prose"))
	fmt.Println(fbgt.Transfer("adventure"))
	// fmt.Println(fbgt.GenreName("prose_su_classics", "ru"))
}
