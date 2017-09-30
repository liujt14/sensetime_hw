# FTP client and server for go
# install
go get github.com/liujt14/sensetime_hw
# Usage
build go resources:
'mkdir myclient&& mv myclient.go ./myclient'
'cd myclient'
'go build'
'cd ..'
'go build'
Open new terminal and run  built files separately 
'./myserver'
'./myclient'
Use following commands on myclient terminal 
# commands
..* user/pass
..* list
..* cwp
..* pwd
..* retr
..* stor 
# To do list
..* open another port as data connection
..* set username and password library
..* improve command enter method
..* build struct and interface plus functions to optimize code structure
