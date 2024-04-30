# DJMocracy

Amateur DJs, meet your new best friend. DJmocracy allows you to invite your 
friends or future party goers to help you create a playlist for your event and
then allows them to vote to put the playlist in order of the most wanted songs.

## Local Dev
In order to run the app currently you must be running Node >= v20. Docker wasn't
working with vite when I originally set this up so to currently run the app do this:

1. open up 3 different terminals one at /, one in /vue-app, and one in /server
2. In / run `docker-compose up`
3. In /server run `go run main.go` to run the backend
4. In /vue-app run `yarn dev` to run the vue app

Sometime soon I'll dockerize the backend to also include a db that it talks to.
When that happens, I'm going to write a Makefile to handle the running of the app
so that it's easier for future travelers to get up and running :)
