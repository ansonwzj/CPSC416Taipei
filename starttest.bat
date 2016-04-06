echo this assumes that your file is in bin, maybe make it better but whatever
set /P filename=Enter FileName:
set /P numberOfClients=Enter the number of clients: 
set trackerAddress="127.0.0.1:8080"
set loggerPort="9000"

set workingDir=%cd%
cd ..\..\..\..\bin 

rm -rf \test
mkdir ..\test
for /L %%a in (1,1,%numberOfClients%) do (
  mkdir ..\test\downloader\%%a
  mkdir ..\test\seed
)

copy %filename% ..\test\seed

start logger.exe 127.0.0.1 %loggerPort%

CPSC416Taipei.exe -createTracker %trackerAddress% -createTorrent %filename% >test.torrent

start CPSC416Taipei.exe -createTracker %trackerAddress% test.torrent 

set origin=%cd% 

for /L %%a in (1,1,%numberOfClients%) do (
  cd ..\test\downloader\%%a
  start ..\..\..\bin\CPSC416Taipei.exe -loggerAddress 127.0.0.1:%loggerPort% -port 0 ..\..\..\bin\test.torrent
  cd %origin%
)

cd ..\test\seed
start ..\..\bin\CPSC416Taipei.exe -loggerAddress 127.0.0.1:%loggerPort% -port 0 ..\..\bin\test.torrent 
cd %origin%
cd %workingDir%
