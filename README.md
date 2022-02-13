# lifelogspd

[![Go Report Card](https://goreportcard.com/badge/github.com/spech66/lifelogspd)](https://goreportcard.com/report/github.com/spech66/lifelogspd) [![GoDoc](https://godoc.org/github.com/spech66/lifelogspd?status.svg)](https://godoc.org/github.com/spech66/lifelogspd)

LifelogSPDaemon - **Journal, weight, Strength training and endurance training tracker**.

This daemon is intend for use in a local network. For securing the access use a proxy like nginx (see below)!

## Status

* **Weight:** :heavy_check_mark:
* **Weight chart:** :heavy_check_mark:
* **Journal:** :heavy_check_mark:
* **Strength training:** :heavy_check_mark:
* **Strength training chart:** :heavy_check_mark:
* **Endurance workout:** :heavy_check_mark:
* **Endurance workout chart:** :heavy_check_mark:

## Screenshots

![Start](https://raw.githubusercontent.com/spech66/lifelogspd/master/_screenshots/s_001_start.png "Start")
![Weight](https://raw.githubusercontent.com/spech66/lifelogspd/master/_screenshots/s_002_weight_01.png "Weight")
![Weight Graph](https://raw.githubusercontent.com/spech66/lifelogspd/master/_screenshots/s_002_weight_02.png "Weight Graph")
![Strength training](https://raw.githubusercontent.com/spech66/lifelogspd/master/_screenshots/s_003_strengthtraining_01.png "Strength training")
![Endurance workout](https://raw.githubusercontent.com/spech66/lifelogspd/master/_screenshots/s_004_enduranceworkout_01.png "Endurance workout")
![Journal](https://raw.githubusercontent.com/spech66/lifelogspd/master/_screenshots/s_005_journal_01.png "Journal")

[More screenshots](https://github.com/spech66/lifelogspd/tree/master/_screenshots)

## Build and run from source

Make sure you have the [Go Programming Language](https://golang.org/) tools set up an ready.

### Linux

Checkout the code to your `GOPATH` directory, build the server and run it.

```bash
go get github.com/spech66/lifelogspd
cd $GOPATH/src/github.com/spech66/lifelogspd
go build
./lifelogspd -config example.config.json
```

### Windows

Checkout the code to your `GOPATH` directory.

```cmd
go get github.com/spech66/lifelogspd
cd %GOPATH%\src\github.com\spech66\lifelogspd
go build
lifelogspd.exe -config example.config.json
```

## nginx reverse proxy with basic auth

This daemon is intend for use in a local network. Make sure that the server port (default 8066) is not exposed to the internet.
In case you want to host this deamon in the internet set up nginx as a revery proxy for it using basic authentication.

```bash
sudo sh -c "echo -n 'admin:' >> /etc/nginx/.htpasswd"
sudo sh -c "openssl passwd -apr1 >> /etc/nginx/.htpasswd"
sudo vi /etc/nginx/sites-enabled/default
```

```nginx
server {
        listen 8666;
        server_name _;
        auth_basic "LifelogSPDaemon";
        auth_basic_user_file /etc/nginx/.htpasswd;

        location / {
                proxy_pass http://localhost:8066;
        }
}
```

## Linux (systemd) Service

Create service file at `/etc/systemd/system/lifelog.service`. Start with `systemctl start lifelog`. Autostart on boot with `systemctl enable lifelog`.

```ini
[Unit]
Description=LifelogSPDeamon
After=network.target

[Service]
Type=simple
ExecStart=/root/go/src/github.com/spech66/lifelogspd/lifelogspd -bind localhost:8066 -config /root/lifelogspd/config.json
Restart=always
# Consider creating a dedicated user here:
User=root
#Environment=
WorkingDirectory=/root/go/src/github.com/spech66/lifelogspd/

[Install]
WantedBy=multi-user.target
```

## Dependencies

* [Go echo web framework](github.com/labstack/echo) - High performance, minimalist Go web framework

## Development

* Update Modules: `go mod tidy`
* Run: `go run . -config example.config.json`
