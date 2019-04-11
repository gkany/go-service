#!/bin/bash

SERVER="go-service"
BASE_DIR="."
INTERVAL=2


function start()
{
        if [ "`pgrep $SERVER -u $UID`" != "" ];then
                echo "$SERVER already running"
                exit 1
        fi
       $BASE_DIR/$SERVER start -c  $BASE_DIR/config.json  &>$BASE_DIR/logs/server.log &

        echo "sleeping..." &&  sleep $INTERVAL

        # check status
        if [ "`pgrep $SERVER -u $UID`" == "" ];then
                echo "$SERVER start failed..."
                exit 1
        fi
}

function status()
{
        if [ "`pgrep $SERVER -u $UID`" != "" ];then
                echo $SERVER is running
        else
                echo $SERVER is not running
        fi
}

function stop()
{
        if [ "`pgrep $SERVER -u $UID`" != "" ];then
                kill -9 `pgrep $SERVER -u $UID`
        fi

        echo "sleeping..." &&  sleep $INTERVAL

        if [ "`pgrep $SERVER -u $UID`" != "" ];then
                echo "$SERVER stop failed..."
                exit 1
        fi
}

function build()
{
    go build -ldflags "-w -s" -o  $BASE_DIR/$SERVER  ../main.go

    echo "sleeping..." &&  sleep $INTERVAL
}


function build_mac_for_windows()
{
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o $BASE_DIR/${SERVER}.exe ../main.go

    echo "sleeping..." &&  sleep $INTERVAL
}

case "$1" in
        'build')
        build
        ;;
        'build2')
        build_mac_for_windows
        ;;
        'start')
        start
        ;;
        'stop')
        stop
        ;;
        'status')
        status
        ;;
        'restart')
        stop && start
        ;;
        *)
        echo "usage: $0 {start config |stop|restart config|status}"
        exit 1
        ;;
esac


