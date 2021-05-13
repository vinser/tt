package main

import (
	"encoding/xml"
	"fmt"
)

type Binary struct {
	Id          string `xml:"id,attr"`
	ContentType string `xml:"content-type,attr"`
	Content     string `xml:",chardata"`
}

func main() {
	str := `<?xml version="1.0" encoding="utf-8"?>
    <binary id="cover.jpg" content-type="image/jpeg">
/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAMCAgMCAgMDAwMEAwMEBQgFBQQEBQoHBwYIDAoM
DAsKCwsNDhIQDQ4RDgsLEBYQERMUFRUVDA8XGBYUGBIUFRT/2wBDAQMEBAUEBQkFBQkUDQsN
FBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBT/wgAR
CAJEAaQDASIAAhEBAxEB/8QAHAAAAQQDAQAAAAAAAAAAAAAABQIDBAYAAQcI/8QAGQEAAwEB
AQAAAAAAAAAAAAAAAAECAwQF/9oADAMBAAIQAxAAAAHrcdbXWSsYxp1vEg0HN1Io7PUqoiam
B0KqUWwxrHssnVZbUhRKa5YazybRI7TvFuVAnB+boUK/Ve7bsYmyaQ1WjY+UtMobkFjleO7Z
2qbXzXp8+23RmyeoXQa2qPLFvVJCK64ge+7oGWZrbEpdQ4AloJIbeOY1Ytadx2ZktLDNqylC
AbsM3kcUMDbs8wm4EKwnBVt1FQnGwUM1V2FcuzUF2Fx6WwfGXDjUi81ebXYBxvXNmGZBZjY8
sISJlRpDabRJ2/6eEV2QnaYwayVUZeU4kE6VpiMUkE6VEacFYYGztxBCcVjLDpD2OiFJS2uL
</binary>
`

	var cp Binary

	err := xml.Unmarshal([]byte(str), &cp)
	fmt.Println(err)
	fmt.Printf(
		`CoverPage ----
  Id: %s
  Content-type: %s
================================
%#v
================================
`,
		cp.Id,
		cp.ContentType,
		cp.Content,
	)

}
