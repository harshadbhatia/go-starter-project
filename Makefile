

install-go:
	wget https://storage.googleapis.com/golang/getgo/installer_linux -O /tmp/installer_linux
	chmod u+x /tmp/installer_linux
	/tmp/installer_linux
	echo "Execute the following execute /tmp/installer_linux and then do source ~/.bash_profile"

start:
	cd src && \
	go mod init starter && \
	go mod tidy

build:
	docker-compose up --build

