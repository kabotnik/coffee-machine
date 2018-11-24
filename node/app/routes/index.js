const nodeRoutes = require('./node_routes');

module.exports = function(app) {
    nodeRoutes(app);
};