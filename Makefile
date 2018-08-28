HEAD = `git rev-parse HEAD`
TIME = `date +%FT%T%z`

BINARY = "go-sdk-example"

default:
	go build -ldflags "-X main.Version=${HEAD} -X main.BuildTime=${TIME}" -o ${BINARY}

clean:
	rm ${BINARY}
