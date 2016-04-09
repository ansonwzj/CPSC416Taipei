#!/bin/bash
#This machine will be a tracker, logger and seeder

# Assuming I'm remote.ugrad.cs.ubc.ca 
servers=( lulu.ugrad.cs.ubc.ca annacis.ugrad.cs.ubc.ca bowen.ugrad.cs.ubc.ca thetis.ugrad.cs.ubc.ca deas.ugrad.cs.ubc.ca )
trackerport=10000
loggerport=10001

#read file
echo "Enter file to seed/track/whatever"
file="fireflies.mp3"
#read file

#kills all running taipeis, Shit way of doing it 
#ps aux | grep [C]PSC416Taipei | awk '{print $2}' | xargs -I pid pkill -9 pid
killall CPSC416Taipei
killall logger

thisip=$(ip addr show em1 | grep inet | awk '{print $2}' | sed s_\/..__g )

./CPSC416Taipei -createTracker $thisip:$trackerport -createTorrent $file > test.torrent
./CPSC416Taipei -createTracker $thisip:$trackerport test.torrent &
./logger $thisip $loggerport &

for server in ${servers[@]}
do
  mkdir -p ~/cs416/test/download/$server
  ssh -t t8e8@$server " 
  cd cs416/test/download/$server;
  nohup ../.././CPSC416Taipei ../../test.torrent &
  " 
done

#for server in ${servers[@]}
#do
#  ssh -t t8e8@$server "
#  cd cs416/test/download/
#done

# ssh -t t8e8@${servers[$i]} "cd cs416/test/download/$i; ../.././CPSC416Taipei ../../test.torrent"
