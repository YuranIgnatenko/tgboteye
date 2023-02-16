# tgboteye
> Descript
```
Detect online/offline mode user
and save history log in/out telegram.
Language ui: Russian
```

***

> Install
```
git clone git@github.com:YuranIgnatenko/tgboteye.git
```
***

> Requirements
```bash
# golang:
go get go-telegram-bot-api/telegram-bot-api
```
```bash
# python3
pip3 install telethon
```

***

> Configure
```bash
# path: ~/tgboteye/config/config_py.json
{
    "api_hash":"8d7sfsdf7sd8yf8g78s7dfy87gdfy8gv",
    "api_id":"998877",
    "delay_sec":0
}
```
```bash
# path: ~/tgboteye/config/config.json
{
"Token": "999992708:AAEiVaaEBxSqZH43o8SGA90ic3ld9lL"
}
```

***

> Launch

Step 1
```bash
# WARNING !!!
cd ~/tgboteye/cmd/
python3 check_inline.py test
# Enter phone-number
# Wait Check-message code
# Enter code
```
Step 2

```bash
cd ~/tgboteye/

go run main.go -c=config/config.json

# or

nohup go run main.go > nohup.out 2>> nohup.out &
tail -f nohup.out

```

***

> Demo screen
> 
![demo](/demo/bot_demo1.png)
![demo](/demo/bot_demo2.png)
![demo](/demo/bot_demo3.png)

***