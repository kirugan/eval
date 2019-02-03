set package=github.com/kirugan/eval
docker run -it -v %cd%:/go/src/%package% golang go test -test.v %package%