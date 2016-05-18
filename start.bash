#!/bin/bash
live=`ps aux | grep movie | grep -v grep`
if [ -z ${live} ]; then
    echo "it is running"
    exit
fi
mv videoInfo.html /home/xxx/webroot/movie/
mv -f pages/* /home/xxx/webroot/movie/pages/
./movie
