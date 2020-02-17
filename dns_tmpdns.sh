#!/bin/bash

#Usage: dns_tmpdns_add   _acme-challenge.www.domain.com   "XKrxpRBosdIKFzxW_CT3KLZNf6q0HG9i01zxXp5CPBs"
dns_tmpdns_add() {
  fulldomain=$1
  txtvalue=$2
  _debug "docker run"
  exists=`docker inspect -f '{{.Config.Cmd}}' tmpdns | awk '{print substr($0, 2,length($0)-2)}'` || true
  docker rm -f tmpdns || true
  docker run -d -p 53:53/udp --name tmpdns binzume/tmpdns $exists "$fulldomain.:txt:$txtvalue"
}

#Usage: fulldomain txtvalue
#Remove the txt record after validation.
dns_tmpdns_rm() {
  fulldomain=$1
  txtvalue=$2
  docker rm -f tmpdns || true
}

