rd /s/q release
md release
go build -o chat.exe
COPY chat.exe release\
COPY favicon.ico release\favicon.ico
XCOPY asset\*.* release\asset\  /s /e
XCOPY view\*.* release\view\  /s /e 