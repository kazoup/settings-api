# Settings API

Settings API is to be used behind Micro API gateway.

API    | ENDPOINT
-------|---------
Public | /v2/settings/public/{create,update,read,delete}
Hostname | /v2/settings/hostname/{create,update,read,delete}
SmbUser | /v2/settings/smbuser/{create,update,read,delete}
CloudStorage | /v2/settings/cloudstorage/{create,update,read,delete}
LDAP | /v2/settings/ldap/{create,update,read,delete}
SMTP | /v2/settings/smtp/{create,update,read,delete}
Logs | /v2/settings/logs/{create,update,read,delete}
Subscription | /v2/settings/subscription/{create,update,read,delete}
DataSource | /v2/settings/datasource/{create,update,read,search,discover,delete}
Policy | /v2/settings/policy/{create,update,read,search,delete}
User | /v2/settings/user/{create,update,read,search,delete}

## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
	```
3. Quick run

	```
	go run main.go
	```
4. Download and start the service

	```shell
	go get github.com/kazoup/settings-api
	settings-api
	```

	OR as a docker container

	```shell
	docker run kazoup/settings-api --registry_address=YOUR_REGISTRY_ADDRESS
	```
