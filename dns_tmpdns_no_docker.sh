#!/bin/bash

#Usage: dns_tmpdns_add   _acme-challenge.www.domain.com   "XKrxpRBosdIKFzxW_CT3KLZNf6q0HG9i01zxXp5CPBs"
dns_tmpdns_no_docker_add() {
  fulldomain=$1
  txtvalue=$2
  if [ ${EUID:-${UID}} = 0 ]; then
    sudo=""
  else
    sudo="sudo"
  fi
  _debug "tmpdns"
  $sudo killall tmpdns || true
  echo "$fulldomain.:txt:$txtvalue" >> records.txt
  $sudo ./tmpdns -p 53 `cat records.txt` &
}

#Usage: fulldomain txtvalue
#Remove the txt record after validation.
dns_tmpdns_no_docker_rm() {
  fulldomain=$1
  txtvalue=$2
  if [ ${EUID:-${UID}} = 0 ]; then
    sudo=""
  else
    sudo="sudo"
  fi
  $sudo killall tmpdns || true
  rm records.txt || true
}

