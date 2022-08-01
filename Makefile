run:
	go run main.go

server:
	go run network/server.go

sys:
	go run system/interface.go


.PHONY: run, server, client