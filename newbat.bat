for /L %%a in (1,1,10) do (
  cd ..\test\downloader\%%a
  echo holyshit > test.pdf
  del test.pdf 
  cd %origin%
)
