module.exports = function(app) {
    app.get('/coffee', (req, res) => {
        res.send('Coffee!');
    });
};