package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

type userInfo struct {
	Description     string `json:"description"`
	FavouritesCount int    `json:"favourites_count"`
	FollowersCount  int    `json:"followers_count"`
	FriendsCount    int    `json:"friends_count"`
	GeoEnabled      bool   `json:"geo_enabled"`
	Location        string `json:"location"` // User defined location
	Name            string `json:"name"`
	Protected       bool   `json:"protected"`
	ScreenName      string `json:"screen_name"`
	TimeZone        string `json:"time_zone"`
}

func searchToUserInfo(user anaconda.User) userInfo {
	return userInfo{
		Description:     user.Description,
		FavouritesCount: user.FavouritesCount,
		FollowersCount:  user.FollowersCount,
		FriendsCount:    user.FriendsCount,
		GeoEnabled:      user.GeoEnabled,
		Location:        user.Location,
		Name:            user.Name,
		Protected:       user.Protected,
		TimeZone:        user.TimeZone,
	}
}

type searchRequest struct {
	Search string `json:"username"`
}

func (db *API) firstPage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "POST" {

		r.ParseForm()
		fmt.Println(r.Form.Get("username"))
		req := searchRequest{
			Search: r.Form.Get("username"),
		}
		// get user name to find here
		searchResult, err := db.api.GetUserSearch(req.Search, nil)
		if err != nil {
			fmt.Println(err)
		}
		userResult := []userInfo{}
		for _, sr := range searchResult {
			userResult = append(userResult, searchToUserInfo(sr))
		}
		t := template.Must(template.ParseFiles("html/twitter_user_footprint.html"))
		err = t.Execute(w, userResult)
		if err != nil {
			fmt.Println(err)
		}
		// return page load with data here
		return
	} else if r.Method == "GET" {
		// return page load here

	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// check url here
		if r.URL.Path != "/" {
			return
		}
		//serve index here
		indexFile, err := ioutil.ReadFile("html/index.html")
		if err != nil {
			http.Error(w, "Error finding index.html", 400)
			fmt.Println(err)
		}
		w.Write(indexFile)
	}

}

func about(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// check url here
		if r.URL.Path != "/about" {
			return
		}
		//serve index here
		aboutFile, err := ioutil.ReadFile("html/about.html")
		if err != nil {
			http.Error(w, "Error finding index.html", 400)
		}
		w.Write(aboutFile)
	}

}

type API struct {
	api *anaconda.TwitterApi
}

func main() {
	CK := os.Getenv("CK")
	CS := os.Getenv("CS")
	AT := os.Getenv("AT")
	ATS := os.Getenv("ATS")
	anaconda.SetConsumerKey(CK)
	anaconda.SetConsumerSecret(CS)
	creds := anaconda.NewTwitterApi(AT, ATS)

	db := API{api: creds}

	http.HandleFunc("/twitter_user_footprint", db.firstPage)
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	CRT := os.Getenv("CRT")
	KEY := os.Getenv("KEY")
	fmt.Println(CRT, KEY)
	for _, i := range []string{"CK", "CS", "AT", "ATS", "CRT", "KEY"} {
		fmt.Println(os.Getenv(i))
	}
	//log.Fatal(http.ListenAndServeTLS(":8443", "ssl/shellcode.in.crt", "ssl/shellcode.in.key", nil))
	log.Fatal(http.ListenAndServeTLS(":8443", CRT, KEY, nil))
}
