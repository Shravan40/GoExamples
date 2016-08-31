In this project list of example given in go tour.

First you need to install go in your machine

# Step 0 -> Update your machine

$ sudo apt-get update

$ sudo apt-get -y upgrade

# Step 1 -> Download and unpack the go package

$ sudo curl -o https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz

$ sudo tar -xvf go1.7.linum-amd64.tar.gz

$ sudo mv go/ /usr/local


# Step 3 -> Set the go path

$ sudo vi ~/.profile

At the end of of the file, add this line

$ export PATH=$PATH:/usr/local/go/bin

$ source ~/.profile

# Step -> Point your project directory to GOPATH

$ export GOPATH=$HOME/path_to_your_project_dir


To see the output of particular file

$ go run file_name.go
