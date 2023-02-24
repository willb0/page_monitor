# page-monitor

i'm writing this repo in go to monitor websites and send notifications of content change using with redis pub sub

you can run page monitor rn if you run a site on localhost:3000 and have docker:

```sh
docker run -p 6379:6379 --name some-redis -d redis
git clone https://github.com/willb0/page_monitor
cd page_monitor
go mod download
go build
./page_monitor -url https://localhost:3000 -refresh_rate 5
```