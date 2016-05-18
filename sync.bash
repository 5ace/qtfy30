#!/bin/bash
live=`ps aux | grep movie | grep -v grep | wc -l`
while [ $live -gt 0 ] 
do
    echo "it is running"
    sleep 100
    live=`ps aux | grep movie | grep -v grep | wc -l`
done
echo $live
cp videoInfo.html /home/xxx/webroot/movie/
cp -f pages/* /home/xxx/webroot/movie/pages/

