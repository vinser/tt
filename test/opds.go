package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {

	rootPage := `<?xml version="1.0" encoding="UTF-8"?>
	<feed xmlns="http://www.w3.org/2005/Atom">
	<title>D-u-u-u-umb</title>
	<updated>2021-03-08T23:04:40-08:00</updated>
	<id>/opds</id>
	<link  rel="search" title="Search Catalog" type="application/atom+xml" href="/opds/search?q={searchTerms}"/>
	<link rel="root" type="application/atom+xml" href="/opds"/>
	<link rel="self" type="application/atom+xml" href="/opds"/>
	<entry>
	  <title>New Titles</title>
	  <id>urn:manybooks:56d70474-2743-1dba-ac46-639fad6af333</id>
	  <updated>2021-03-08T23:04:40-08:00</updated>
	  <content type="text">Recently Added books</content>
	  <link type="application/atom+xml" href="/opds/new_titles"/>
	</entry>
	<entry>
	  <title>Authors</title>
	  <id>urn:manybooks:364fe68a-3bab-8921-ad8d-f271587da312</id>
	  <updated>2021-03-08T23:04:40-08:00</updated>
	  <content type="text">Alphabetical listing</content>
	  <link type="application/atom+xml" href="/opds/authors"/>
	</entry>
	<entry>
	  <title>Genres</title>
	  <id>urn:manybooks:1d946c89-6153-307b-d07b-28a21f76c379</id>
	  <updated>2021-03-08T23:04:40-08:00</updated>
	  <content type="text">Titles grouped by category</content>
	  <link type="application/atom+xml" href="/opds/genres"/>
	</entry>
  </feed>`

	// 	newTitlesPage := `<?xml version="1.0" encoding="utf-8"?><feed xmlns="http://www.w3.org/2005/Atom">
	// <title>D-u-u-u-umb</title>
	//   <updated>2021-03-08T23:04:40-08:00</updated>
	//   <author>
	// 	<name>192.168.0.20:8080</name>
	// 	<uri>http://192.168.0.20:8080/</uri>
	// 	<email>info@192.168.0.20:8080</email>
	//   </author>
	//   <subtitle>
	//     Recently Added books
	//   </subtitle>
	//   <id>urn:manybooks:d546f861-bc57-7ace-b943-a10b377ba8ba</id>
	//   <link  rel="search" title="Search Catalog" type="application/atom+xml" href="/opds/search?q={searchTerms}"/>
	//   <link rel="root" type="application/atom+xml" href="/opds"/>
	//   <link rel="self" type="application/atom+xml" href="/opds/new_titles"/>
	//   <entry>
	// 	<title>Authors</title>
	// 	<id>urn:manybooks:364fe68a-3bab-8921-ad8d-f271587da312</id>
	// 	<updated>2021-03-08T23:04:40-08:00</updated>
	// 	<content type="text">Alphabetical listing</content>
	// 	<link type="application/atom+xml" href="/opds/authors"/>
	//   </entry>
	//   <entry>
	// 	<title>Titles</title>
	// 	<id>urn:manybooks:80daf494-787c-385a-b2ae-05e36397e3d6</id>
	// 	<updated>2021-03-08T23:04:40-08:00</updated>
	// 	<content type="text">All titles, listed alphabetically</content>
	// 	<link type="application/atom+xml" href="/opds/titles"/>
	//   </entry>  <entry>
	// 	<title>Genres</title>
	// 	<id>urn:manybooks:1d946c89-6153-307b-d07b-28a21f76c379</id>
	// 	<updated>2021-03-08T23:04:40-08:00</updated>
	// 	<content type="text">Titles grouped by category</content>
	// 	<link type="application/atom+xml" href="/opds/genres"/>
	//   </entry>
	// </feed>`

	opdsRootHandler := func(w http.ResponseWriter, req *http.Request) {
		qu, _ := url.QueryUnescape(req.URL.String())
		fmt.Printf("DUMB->URL: [%s]\n", qu)
		w.Header().Set("content-type", "application/atom+xml")
		io.WriteString(w, rootPage)
	}

	// opdsNewTitlesHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	qu, _ := url.QueryUnescape(req.URL.String())
	// 	fmt.Printf("DUMB->URL: [%s]\n", qu)
	// 	w.Header().Set("content-type", "application/atom+xml")
	// 	io.WriteString(w, newTitlesPage)
	// }

	http.HandleFunc("/opds", opdsRootHandler)
	// http.HandleFunc("/opds/new_titles", opdsNewTitlesHandler)
	// http.HandleFunc("/opds/authors", opdsRootHandler)
	// http.HandleFunc("/opds/titles", opdsRootHandler)
	// http.HandleFunc("/opds/genres", opdsRootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
