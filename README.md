# TmpDNS

テンポラリなDNSサーバー．
設定ファイルなどは存在せず，起動時のコマンドライン引数のみで完結します．

一瞬だけDNSサーバを立てたくなったときに便利です．

ACMEのDNS認証用にも使えます．
(dns_tmpdns.sh がacme.shから使えます)

# Usage

```console
go get -d
CGO_ENABLED=0 go build
./tmpdns -p 53 "hoge.example.com.:txt:hello! world" "fuga.example.com.:a:192.168.0.1"
./tmpdns -p 53 -z example.com. "hoge:txt:hello! world" "fuga:a:192.168.0.1"
```
(FQDNを示す `.` の書き忘れに注意)

## Docker image

```console
dokcer pull binzume/tmpdns
docker run --rm -p 53:53/udp --name tmpdns binzume/tmpdns -z example.com. "hoge:txt:hello! world" "fuga:a:192.168.0.1"
```

## Flags

- -p: port (default:53)
- -z: zone (default:.)

# License

TOOD:

MIT License
