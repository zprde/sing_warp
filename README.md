# sing_warp
Integrate singbox and cloudflare warp.

## requirements
- Docker
- Docker-compose

## How to start
```shell
docker-compose build
docker-compose up -d sing-warp
```

Then the `0.0.0.0:2080` can be a http/https/sock5 proxy port.

## How to uninstall it
```shell
docker-compose down --rmi all
```

Reference Repos:
- [wgcf](https://github.com/ViRb3/wgcf)
- [sing-box](https://github.com/SagerNet/sing-box)
- [warp-yg](https://github.com/yonggekkk/warp-yg)
