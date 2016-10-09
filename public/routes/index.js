var express = require('express');
var router = express.Router();

/* Present the home page. */
router.get('/', function(req, res, next) {

	res.render('index', { title: 'MyBigAssFootprint' } );

});

module.exports = router;
