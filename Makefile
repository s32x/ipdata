init:
	-rm -rf ./vendor go.mod go.sum
	GO111MODULE=on go mod init

deps:
	-rm -rf ./vendor go.mod go.sum
	go clean --modcache
	go mod init
	go mod tidy
	go mod vendor

db:
	-rm -rf ./db
	mkdir db
	wget -O ./db/city.tar.gz http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz
	wget -O ./db/asn.tar.gz http://geolite.maxmind.com/download/geoip/database/GeoLite2-ASN.tar.gz

test:
	go test ./...

deploy: deps db test
	up prune -s production -r 2
	-up stack plan
	-up stack apply
	up deploy production