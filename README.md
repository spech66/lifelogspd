# lifelogspd

LifelogSPDaemon - Data collection/editor related to [LifelogSP](https://github.com/spech66/lifelogsp).

This daemon is intend for use in a local network. For securing the access use a proxy like nginx (see below)!

## Status

* **Weight:** Mostly working except for deletion of data :chart_with_upwards_trend:
* **Weight Chart:** Working :heavy_check_mark: (Lines for good/bad bmi/weight might be added in the future)
* **Journal:** Work in progress :construction:
* **Strength training:** Work in progress :construction:
* **Endurance workout:** Work in progress :construction:

![Start](https://raw.githubusercontent.com/spech66/lifelogspd/master/_screenshots/001_start.png "Start")
![Weight](https://raw.githubusercontent.com/spech66/lifelogspd/master/_screenshots/002_weight_01.png "Weight")
![Weight Graph](https://raw.githubusercontent.com/spech66/lifelogspd/master/_screenshots/002_weight_02.png "Weight Graph")


## nginx reverse proxy with basic auth

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

## Dependencies

* [Go echo web framework](github.com/labstack/echo) - High performance, minimalist Go web framework
