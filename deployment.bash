#!/bin/bash
#This machine will be a tracker, logger and seeder

# Assuming I'm remote.ugrad.cs.ubc.ca 
#servers=( lulu.ugrad.cs.ubc.ca annacis.ugrad.cs.ubc.ca bowen.ugrad.cs.ubc.ca thetis.ugrad.cs.ubc.ca deas.ugrad.cs.ubc.ca lin01.ugrad.cs.ubc.ca lin02.ugrad.cs.ubc.ca lin03.ugrad.cs.ubc.ca lin04.ugrad.cs.ubc.ca lin06.ugrad.cs.ubc.ca lin07.ugrad.cs.ubc.ca lin08.ugrad.cs.ubc.ca lin09.ugrad.cs.ubc.ca lin10.ugrad.cs.ubc.ca lin11.ugrad.cs.ubc.ca lin12.ugrad.cs.ubc.ca lin13.ugrad.cs.ubc.ca lin14.ugrad.cs.ubc.ca lin15.ugrad.cs.ubc.ca lin16.ugrad.cs.ubc.ca lin17.ugrad.cs.ubc.ca lin18.ugrad.cs.ubc.ca lin19.ugrad.cs.ubc.ca lin20.ugrad.cs.ubc.ca lin21.ugrad.cs.ubc.ca lin22.ugrad.cs.ubc.ca lin23.ugrad.cs.ubc.ca lin24.ugrad.cs.ubc.ca lin25.ugrad.cs.ubc.ca)

policies=( Fair Classic Never Random ) 

for policy in ${policies[@]}
do
  for i in $(seq 1 4)
  do
    servers=( lulu.ugrad.cs.ubc.ca annacis.ugrad.cs.ubc.ca bowen.ugrad.cs.ubc.ca thetis.ugrad.cs.ubc.ca deas.ugrad.cs.ubc.ca lin02.ugrad.cs.ubc.ca lin03.ugrad.cs.ubc.ca lin04.ugrad.cs.ubc.ca lin06.ugrad.cs.ubc.ca lin07.ugrad.cs.ubc.ca lin08.ugrad.cs.ubc.ca lin09.ugrad.cs.ubc.ca )

    rm -rf download
    rm ratiof.csv upload.csv download.csv deviation.csv

    #servers=( lulu.ugrad.cs.ubc.ca annacis.ugrad.cs.ubc.ca bowen.ugrad.cs.ubc.ca thetis.ugrad.cs.ubc.ca deas.ugrad.cs.ubc.ca )
    #servers=( annacis.ugrad.cs.ubc.ca )
    trackerport=10000
    loggerport=10001

    echo Which policy to use?
    #read policy 
    echo using: $policy

    echo "Enter file to seed/track/whatever"
    file="book.pdf"
    #read file

    #kills all running taipeis, Shit way of doing it 
    #ps aux | grep [C]PSC416Taipei | awk '{print $2}' | xargs -I pid pkill -9 pid
    killall CPSC416Taipei
    killall logger

    thisip=$(ip addr show em1 | grep inet | awk '{print $2}' | sed s_\/..__g | head -1 )
    echo Got this ip: $thisip

    ./CPSC416Taipei -createTracker $thisip:$trackerport -createTorrent $file > test.torrent
    ./CPSC416Taipei -createTracker $thisip:$trackerport test.torrent &
    ./logger $thisip $loggerport &

    sleep 2
    ./CPSC416Taipei -port 0 -useChokingAlgorithm $policy -loggerAddress $thisip:$loggerport ~/cs416/test/test.torrent &

    seederPID=$!

    for server in ${servers[@]}
    do
      mkdir -p ~/cs416/test/download/$server
      ssh t8e8@$server "~/cs416/test/CPSC416Taipei -useChokingAlgorithm $policy -port 0 -fileDir ~/cs416/test/download/$server -loggerAddress $thisip:$loggerport ~/cs416/test/test.torrent > /dev/null 2> /dev/null < /dev/null &"
    done

    sleep 50 
    ./killall.bash

    mv ratiof.csv ratio$policy$i.csv
    mv upload.csv upload$policy$i.csv
    mv download.csv download$policy$i.csv
    mv deviation.csv deviation$policy$i.csv
  done
done
