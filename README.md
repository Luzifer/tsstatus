[![Go Report Card](https://goreportcard.com/badge/github.com/Luzifer/tsstatus)](https://goreportcard.com/report/github.com/Luzifer/tsstatus)
![](https://badges.fyi/github/license/Luzifer/tsstatus)
![](https://badges.fyi/github/downloads/Luzifer/tsstatus)
![](https://badges.fyi/github/latest-release/Luzifer/tsstatus)
![](https://knut.in/project-status/tsstatus)

# Luzifer / tsstatus

`tsstatus` is a small utility to expose a status of a TeamSpeak3 server.

This can be used to have a monitoring for the server (HTTP 200 vs. HTTP 500) and to retrieve a list of users being present in a channel

```console
# curl -sSf http://localhost:3000/status | jq .
{
  "info": {
    "server": {
      "clients_online": 2,
      "host_button_gfxurl": "https://knut.cc/permanent/1a79a0/luzifer_220px.svg.png",
      "host_button_url": "https://luzifer.io/",
      "max_clients": 32,
      "name": "TS @ luzifer.io",
      "port": 9987,
      "status": "online",
      "uptime": 8219,
      "version": "3.10.2 [Build: 1574239171]",
      "welcome_message": "Welcome to TeamSpeak on luzifer.io!"
    },
    "channels": [
      {
        "id": 1,
        "name": "Lobby",
        "clients": [
          {
            "away": false,
            "away_message": "",
            "nickname": "Luzifer"
          }
        ]
      },
      {
        "id": 2,
        "name": "Game I",
        "clients": null
      },
      {
        "id": 3,
        "name": "Game II",
        "clients": null
      },
      {
        "id": 4,
        "name": "Game III",
        "clients": null
      },
      {
        "id": 9,
        "name": "AFK",
        "clients": null
      }
    ]
  }
}
```
