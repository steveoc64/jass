all: sassgen templegen app-assets appjs sv run
	echo all done


build: sassgen templegen app-assets appjs sv 

get: 
	go get -u honnef.co/go/simple/cmd/gosimple
	go get -u github.com/gopherjs/gopherjs
	go get -u github.com/gopherjs/websocket
	go get -u github.com/go-humble/temple
	go get -u github.com/go-humble/form
	go get -u github.com/go-humble/router
	go get -u github.com/go-humble/locstor
	go get -u github.com/steveoc64/formulate
	go get -u github.com/steveoc64/godev/echocors
	go get -u github.com/steveoc64/godev/sms
	go get -u github.com/steveoc64/godev/smt
	go get -u github.com/steveoc64/godev/db
	go get -u github.com/steveoc64/godev/config
	go get -u honnef.co/go/simple/cmd/gosimple
	go get -u github.com/rs/cors
	go get -u gopkg.in/gomail.v2
	go get -u github.com/labstack/echo
	go get -u github.com/labstack/echo/middleware
	go get -u github.com/lib/pq
	go get -u gopkg.in/mgutz/dat.v1/sqlx-runner
	go get -u github.com/nfnt/resize
	go get -u gopkg.in/gomail.v2
	go get -u github.com/logpacker/PayPal-Go-SDK
	mkdir -p scripts
	mkdir -p backup

help: 
	# sassgen    - make SASS files
	# templegen  - make Templates
	# app-assets - make Asset copy to dist	
	# appjs      - make Frontend app
	# sv         - make Server
	# run        - run  Server

clean:	
	# Delete existing build
	@mplayer -quiet ../audio/trash-empty.oga 2> /dev/null > /dev/null &
	rm -rf dist

sassgen: dist/public/css/jass.css

dist/public/css/jass.css: scss/*
	@mplayer -quiet ../audio/attention.oga 2> /dev/null > /dev/null
	@mkdir -p dist/public/css
	cd scss && node-sass --output-style compressed app.sass ../dist/public/css/jass.css
	cd scss && node-sass app.sass ../dist/public/css/jass.debug.css

templegen: app/template.go 

app/template.go: templates/*.tmpl 	
	@mplayer -quiet ../audio/attention.oga 2> /dev/null > /dev/null
	temple build templates app/template.go --package main

app-assets: dist/assets.log dist/config.json

dist/config.json: server/config.json
	cp server/config.json dist	

cert:
	mkdir -p cert
	openssl genrsa -out cert/jass.key 2048
	openssl req -new -x509 -key cert/jass.key -out cert/jass.pem -days 3650

dist/assets.log: assets/*.html assets/img/*  assets/fonts/* assets/*.webmanifest assets/img/models/* assets/img/items/* assets/css/*
	@mplayer -quiet ../audio/attention.oga 2> /dev/null > /dev/null
	@mkdir -p dist/public/css dist/public/font dist/public/js
	cp assets/*.html dist/public
	cp assets/*.webmanifest dist/public
	cp -R assets/img dist/public
	cp -R assets/fonts dist/public
	cp -R assets/css dist/public
	@date > dist/assets.log

appjs: dist/public/jass.js

dist/public/jass.js: app/*.go shared/*.go makefile app/*.inc.js
	@mplayer -quiet ../audio/frontend-compile.ogg 2> /dev/null > /dev/null &
	@mkdir -p dist/public/js
	#cd app && gosimple
	# @echo -n Before :
	# @ls -l dist/public/jass.js
	GOOS=linux gopherjs build app/*.go app/*.inc.js -o dist/public/jass.js -m
	gopher-count dist/public/jass.js | sort -n
	@ls -l dist/public/jass.js

remake: 
	@mplayer -quiet ../audio/server-compile.oga 2> /dev/null > /dev/null &
	rm -f dist/jass-server
	@gosimple server
	go build -o dist/jass-server server/*.go
	@ls -l dist/jass-server

sv: dist/jass-server 

dist/jass-server: server/*.go shared/*.go
	@mplayer -quiet ../audio/server-compile.oga 2> /dev/null > /dev/null &
	cd server && gosimple
	go build -o dist/jass-server server/*.go
	@ls -l dist/jass-server
	cp cert/jass.key cert/jass.pem dist

run: 
	./terminate
	@mplayer -quiet ../audio/running.oga 2> /dev/null > /dev/null &
	@cd dist && ./jass-server

testsvr: sv
	./terminate-test
	cp -Rv dist/* ~/jass/test
	rm ~/jass/test/config.json
	cp server/config.json.test ~/jass/test/config.json
	cd ~/jass/test && mv jass-server jass-test-server &&  nohup ./jass-test-server &

install: sv
	./terminate
	cp -Rv dist/* ~/jass/current
	cd ~/jass/current && nohup ./jass-server &
	# tail -f -n 200 ~/logs/jass/* ~/jass/current/nohup.out

tail:
	tail -f -n 200 ~/logs/jass/* ~/jass/current/nohup.out

data:
	pg_dump jass > database/jass.sql
	scp -P 446 database/jass.sql freebsd@bsd:/home/freebsd/jass.sql

loadtempdata:
	./terminate
	dropdb jass
	createdb jass
	psql jass < ~/jass.sql
	cd ~/jass/current && nohup ./jass-server &

loaddata:
	./terminate
	dropdb jass
	createdb jass
	psql jass < database/jass.sql

