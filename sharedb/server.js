var http = require('http');
var express = require('express');
var ShareDB = require('sharedb');
// var redis = require("redis");
var WebSocket = require('ws');
var WebSocketJSONStream = require('@teamwork/websocket-json-stream');


const db = require('sharedb-mongo')('mongodb://mongo:27017/local');

// var client = redis.createClient('redis://user@redis:6379');

// var redisPubsub = require('sharedb-redis-pubsub')(client); // Redis client being an existing redis client connection

var backend = new ShareDB({
  db: db,  // db would be your mongo db or other storage location
  // pubsub: redisPubsub
});

var backend = new ShareDB();

createDoc(startServer);





// Create initial document then fire callback
function createDoc(callback) {
  var connection = backend.connect();
  var doc = connection.get('examples', 'counter');
  doc.fetch(function(err) {
    if (err) throw err;
    if (doc.type === null) {
      doc.create({numClicks: 0}, callback);
      return;
    }
    callback();
  });
}

function startServer() {
  // Create a web server to serve files and listen to WebSocket connections
  var app = express();
  app.use(express.static('static'));
  var server = http.createServer(app);

  // Connect any incoming WebSocket connection to ShareDB
  var wss = new WebSocket.Server({server: server});
  wss.on('connection', function(ws, req) {
    var stream = new WebSocketJSONStream(ws);
    backend.listen(stream);
  });

  server.listen(8080);
  console.log('Listening on http://localhost:8080');
}
