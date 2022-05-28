export PATH=$PATH:/usr/local/go/bin
kitex -type protobuf -module userser -service userdemo -I ../../idl/ ../../idl/user.proto