# lifelogspd

LifelogSPDaemon - Data collection/editor related to [LifelogSP](https://github.com/spech66/lifelogsp).

This daemon is intend for use in a local network. For securing the access use a proxy like nginx (see below)!

## Status

* **Weight:** Working :heavy_check_mark:
* **Weight Chart:** Working :heavy_check_mark:
* **Journal:** Working :heavy_check_mark:
* **Strength training:** Work in progress :chart_with_upwards_trend:
* **Endurance workout:** Work in progress :construction:

### Improvements

* Weight Graph: Baselines for good/bad bmi/weight
* Graphs for training

## Screenshots

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
