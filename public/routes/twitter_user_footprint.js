var express = require('express');
var router = express.Router();

let fetch = require("node-fetch")

 /* Post username to get the possible matches. */
router.post('/', function(req, res, next) {
  var user_request_json_object = {'username': req.body.username};

  console.log("user_object = " + JSON.stringify(user_request_json_object));

  fetch('https://mybigassfootprint.com/twitter_user_footprint', {
    method: 'POST',
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify(user_request_json_object)

  }).then(response => {

    return response.json();

  }).then(json => {

    console.log(JSON.stringify(json));

    json.forEach(
      function logArrayElements(element, index, array) {
      console.log('a[' + index + '] = ' + element.name);
    })

    console.log(json);
    res.render('twitter_user_footprint', {

      twitter_user_footprint_title: "List of Possible Matches",
      jsonObjectResponse: json

     });
  }).catch(err => {


    console.log("error = " + err);

    res.render('index', {

      sorry_message: "Some error occured. Please try again."

    });

  });

});

module.exports = router;


/*

    description: '',
    favourites_count: 2,
    followers_count: 6,
    friends_count: 17,
    geo_enabled: false,
    location: 'canton',
    name: 'Brian',
    protected: false,
    screen_name: '',
    time_zone: '' 

*/