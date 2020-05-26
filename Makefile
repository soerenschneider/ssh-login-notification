tests: unittest integrationtest

unittest:
	go test ./... -coverprofile cover.out

mountebank: clean
	podman run -d --rm --name sshnotmountebank -v ${PWD}/mountebank:/mountebank -p 2525:2525 -p 8080:8080 registry.gitlab.com/soerenschneider/mountebank-docker --pidfile /tmp/mb.pid --logfile /tmp/mb.log --allowInjection --configfile=/mountebank/ip-api.json

integrationtest: 
	go test ./... -tags=integration

containerized: mountebank
	sleep 1
	go test ./... -tags=integration
	podman rm -f sshnotmountebank || true

clean:
	podman rm -f sshnotmountebank || true

build:
	go build

crosscompile:
	if [ ! -d build ]; then mkdir build; fi
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/ssh-login-notification.amd64 ./main.go
	env GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 go build -o build/ssh-login-notification.arm7 ./main.go
	env GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=0 go build -o build/ssh-login-notification.arm5 ./main.go
