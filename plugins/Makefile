SHELL := /bin/bash
.SILENT:
run:	
	if [ -a miplugin.so ]; then \
		rm miplugin.so; \
	fi	
	go build -buildmode=plugin -o miplugin.so ./plugin/plugin.go
	go run ./main.go
	if [ -a miplugin.so ]; then \
		rm miplugin.so; \
	fi	

	