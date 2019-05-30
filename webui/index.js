//https://www.compose.com/articles/redis-pubsub-node-and-socket-io/

var app = require('express')();
var http = require('http').Server(app);
var io = require('socket.io')(http);
var redis = require('redis');

//! incase of docker need to specify redis server separately
//! running REDIS in local machine, thats why need to specify the IP
var client1 = redis.createClient(6379, "172.25.16.126");

app.get('/', function(req, res) {
   res.sendfile('index.html');
});

io.on('connection', onConnection);

http.listen(3000, function() {
   console.log('listening on *:3000');
});

function onConnection(socket){
  console.log('A user connected');

  client1.on('message', function(chan, msg) {
    console.log(msg);
    socket.emit('redisPublishEvent', {description : msg});
  });

  socket.on('disconnect', function(){
    console.log('A user disconnected')
  })
}

client1.subscribe('hashChannel');
