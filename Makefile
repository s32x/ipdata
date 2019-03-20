deps:
	-rm -rf vendor
	-rm -rf db
	-rm -f go.mod
	-rm -f go.sum
	go clean
	GO111MODULE=on go mod init
	GO111MODULE=on go mod vendor
	mkdir db
	wget -O ./db/city.tar.gz http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz
	wget -O ./db/asn.tar.gz http://geolite.maxmind.com/download/geoip/database/GeoLite2-ASN.tar.gz
test:
	go test ./...
install:
	make deps
	go install