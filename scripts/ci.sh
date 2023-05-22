# git clone the application code repository
git 
cd /tmp/sample-web-app

#RUN test with coverage in go environment

go test -cover

# Build docker image 
export $COMMIT_HASH=
