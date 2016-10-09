package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	//"bufio"

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

type abe struct {
	Search        string `json:"username"`
}

func (db *API) firstPage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "POST" {
		//fmt.Println(bufio.NewReader(r.Body).ReadString('\n'))
		var rep abe
		if r.Body == nil {
			http.Error(w, "Please send a request body", 402)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&rep)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(rep)

		// get user name to find here
		searchResult, err := db.api.GetUserSearch(rep.Search, nil)
		if err != nil {
			fmt.Println(err)
		}
		userResult := []userInfo{}
		for _, sr := range searchResult {
			userResult = append(userResult, searchToUserInfo(sr))
		}
		json.NewEncoder(w).Encode(userResult)
		return
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

	CRT := os.Getenv("CRT")
	KEY := os.Getenv("KEY")
	for _, i := range []string{"CK", "CS", "AT", "ATS","CRT", "KEY"} {
		fmt.Println(os.Getenv(i))
	}
	log.Fatal(http.ListenAndServeTLS(":443", CRT, KEY,  nil))
	//log.Fatal(http.ListenAndServeTLS(":8080", "ssl/shellcode.in.crt", "ssl/shellcode.in.key",  nil))
}
