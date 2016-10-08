var express = require('express');
var router = express.Router();

let fetch = require("node-fetch")

/* GET home page. */
router.get('/', function(req, res, next) {
	fetch('https://shellcode.in:8080/first', {
	  method: 'GET',
	  headers: {'Content-Type': 'application/json'},
	  body: '{}'
	}).then(response => {
		console.log(response)
	  return response.json();
	}).then(json => {

		json.forEach(
			function logArrayElements(element, index, array) {
  		console.log('a[' + index + '] = ' + element.name);
		})

		console.log(json)
    res.render('index', { 

    	title: "Tweets:",
    	json_object: json

     });
	}).catch(err => {

		res.render('index', { title: "Some error." } )

		console.log(err);

	})
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