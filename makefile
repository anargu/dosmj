
run_test:
	export $$(cat test.env | xargs) && go test -v ./

test:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes