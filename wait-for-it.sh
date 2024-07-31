#!/usr/bin/env bash
# Use this script to test if a given TCP host/port are available

# The MIT License (MIT)
# Copyright (c) 2014 Michael Stenta

# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:

# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.

# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.

set -e

TIMEOUT=15
QUIET=0
HOST=""
PORT=""

usage()
{
    echo "Usage: $0 host:port [-t timeout] [-- command args]"
    echo " -q | --quiet                        Do not output any status messages"
    echo " -t TIMEOUT | --timeout=timeout      Timeout in seconds, zero for no timeout"
    exit 1
}

wait_for()
{
    if [ "$TIMEOUT" -gt 0 ]; then
        echo "Waiting $TIMEOUT seconds for $HOST:$PORT"
    else
        echo "Waiting for $HOST:$PORT without a timeout"
    fi

    start_ts=$(date +%s)
    while :
    do
        if [ $QUIET -eq 0 ]; then
            (echo > /dev/tcp/$HOST/$PORT) >/dev/null 2>&1
        else
            (echo > /dev/tcp/$HOST/$PORT) >/dev/null 2>&1
        fi

        result=$?
        if [ $result -eq 0 ] ; then
            end_ts=$(date +%s)
            echo "$HOST:$PORT is available after $((end_ts - start_ts)) seconds"
            break
        fi
        sleep 1
    done
    return $result
}

while [ $# -gt 0 ]
do
    case "$1" in
        *:* )
        HOST=$(printf "%s\n" "$1"| cut -d : -f 1)
        PORT=$(printf "%s\n" "$1"| cut -d : -f 2)
        shift 1
        ;;
        -q | --quiet)
        QUIET=1
        shift 1
        ;;
        -t)
        TIMEOUT="$2"
        shift 2
        ;;
        --timeout=*)
        TIMEOUT="${1#*=}"
        shift 1
        ;;
        --)
        shift
        break
        ;;
        *)
        usage
        ;;
    esac
done

if [ "$HOST" = "" -o "$PORT" = "" ]; then
    usage
fi

wait_for
exec "$@"