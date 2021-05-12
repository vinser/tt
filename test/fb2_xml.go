package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type FB2 struct {
	*TitleInfo
}

func main() {
	xmlfile := "/books/578379.fb2"
	xmlStream, err := os.Open(xmlfile)
	if err != nil {
		log.Println("Failed to open XML file")
		return
	}

	fb, err := NewFB2(xmlStream)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	xmlStream.Close()
	fmt.Print(fb)

	xmlStream, err = os.Open(xmlfile)
	cp, err := GetCoverPageBinary(fb.CoverPage.Href, xmlStream)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	xmlStream.Close()
	fmt.Print(cp)
}

func NewFB2(xmlStream *os.File) (*FB2, error) {
	decoder := xml.NewDecoder(xmlStream)
	decoder.CharsetReader = charsetReader
	fb := &FB2{}
TokenLoop:
	for {
		t, err := decoder.Token()
		if t == nil {
			return nil, err
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "title-info" {
				decoder.DecodeElement(fb, &se)
				break TokenLoop
			}
		default:
		}
	}
	return fb, nil
}

func (fb *FB2) String() string {
	return "" + fmt.Sprint(
		"=========FB2===================\n",
		fmt.Sprintf("Authors:    %#v\n", fb.Authors),
		fmt.Sprintf("BookTitle:  %#v\n", fb.BookTitle),
		fmt.Sprintf("Gengre:     %#v\n", fb.Gengres),
		fmt.Sprintf("Annotation: %#v\n", fb.Annotation),
		fmt.Sprintf("Date:       %#v\n", fb.Date),
		fmt.Sprintf("Year:       %#v\n", fb.Year),
		fmt.Sprintf("Lang:       %#v\n", fb.Lang),
		fmt.Sprintf("Serie:      %#v\n", fb.Serie),
		fmt.Sprintf("CoverPage:  %#v\n", fb.CoverPage),
		"===============================\n",
	)
}

type TitleInfo struct {
	Authors    []Author   `xml:"author"`
	BookTitle  string     `xml:"book-title"`
	Gengres    []string   `xml:"genre"`
	Annotation Annotation `xml:"annotation"`
	Date       string     `xml:"date"`
	Year       string     `xml:"year"`
	Lang       string     `xml:"lang"`
	Serie      Serie      `xml:"sequence"`
	CoverPage  Image      `xml:"coverpage>image"`
}

type Author struct {
	FirstName  string `xml:"first-name"`
	MiddleName string `xml:"middle-name"`
	LastName   string `xml:"last-name"`
}

type Annotation struct {
	Text string `xml:",innerxml"`
}

type Serie struct {
	Name   string `xml:"name,attr"`
	Number int    `xml:"number,attr"`
}

type CoverPage struct {
	*Image `xml:"image"`
}

type Image struct {
	Href string `xml:"http://www.w3.org/1999/xlink href,attr"`
}

type Binary struct {
	Id          string `xml:"id,attr"`
	ContentType string `xml:"content-type,attr"`
	Content     []byte `xml:",chardata"`
}

func GetCoverPageBinary(coverLink string, xmlStream *os.File) (*Binary, error) {
	decoder := xml.NewDecoder(xmlStream)
	decoder.CharsetReader = charsetReader
	b := &Binary{}
	coverLink = strings.TrimPrefix(coverLink, "#")
TokenLoop:
	for {
		t, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "binary" && se.Attr[0].Value == coverLink {
				decoder.DecodeElement(b, &se)
				break TokenLoop
			}
		default:
		}
	}
	if b == nil {
		return nil, errors.New("FB2 has no Cover Page")
	}
	return b, nil

}

func (b *Binary) String() string {
	return fmt.Sprintf(
		`CoverPage ----
  Id: %s
  Content-type: %s
================================
%#v
===========(100)================
`, b.Id, b.ContentType, b.Content[:99])
}

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(charset) {
	case "windows-1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	default:
		return nil, fmt.Errorf("unknown charset: %s", charset)
	}
}
