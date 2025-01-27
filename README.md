# fasms
financial assistance scheme management system

## Prerequisites

### macOS

* `brew` to install tools.
* Install `sqitch`.

```sh
brew tap sqitchers/sqitch
brew install sqitch --with-postgres-support

### Other Platforms
* See [official distributions](https://sqitch.org/download/).

## Configure
sqitch config --user user.name ${Your_Full_Name_on_GitLab}
sqitch config --user user.email ${your-email-address@.com}

## Other repo
git clone https://github.com/lftzzzzfeng/fasms-db.git

## Install fasms db
sqitch deploy
```

## Running
* Setup local postgres db with the following loaded:
  * Latest fasms schema, refer to above link.

### macOS
* Update `dev.yaml` config file base on your environment.
* Run `make build`.
* Run `./fasms`.

### docker
* update `host` to `host.docker.internal` from `dev.yaml`.
* build image `docker build -t fasms .`.
* run app `docker run -d -p 8080:8080 fasms`
* check app logs `docker logs [container_id]`