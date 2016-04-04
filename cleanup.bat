taskkill /F /IM Taipei-Torrent.exe /T

timeout 5

set origin=%cd% 

for /L %%a in (1,1,10) do (
  cd ..\test\downloader\%%a
  rm -f * 
  cd %origin%
)

cd ..\test\seed
rm -f * 
cd %origin%
