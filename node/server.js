const express = require('express');

const app = express();

const port = 9090;

require('./app/routes')(app, {});

app.listen(port, () => {
    console.log('Ready to make coffee on ' + port);
});