tests: unittest integrationtest

unittest:
	go test ./... -coverprofile cover.out

mountebank: clean
	podman run -d --rm --name sshnotmountebank -v ${PWD}/mountebank:/mountebank -p 2525:2525 -p 8080:8080 registry.gitlab.com/soerenschneider/mountebank-docker --pidfile /tmp/mb.pid --logfile /tmp/mb.log --allowInjection --configfile=/mountebank/ip-api.json

integrationtest: mountebank
	sleep 1
	go test ./... -tags=integration
	podman rm -f sshnotmountebank || true

clean:
	podman rm -f sshnotmountebank || true
