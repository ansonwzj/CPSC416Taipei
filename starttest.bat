start Taipei-Torrent.exe -createTracker 127.0.0.1:8080 test.torrent
set origin=%cd% 

for /L %%a in (1,1,5) do (
  cd ..\test\downloader\%%a
  start ..\..\..\bin\Taipei-Torrent.exe -port 0 ..\..\..\bin\test.torrent
  cd %origin%
)

cd ..\test\seed
start ..\..\bin\Taipei-Torrent.exe -port 0 ..\..\bin\test.torrent 
cd %origin%
