import os
import sys
sys.path.append(os.getcwd() + '/pb')

import grpc
from pb import user_pb2, user_pb2_grpc

channel = grpc.insecure_channel('localhost:50005')
stub = user_pb2_grpc.UserStub(channel)

req = user_pb2.Request(ID=1)
req.ID = 1

print(stub.Info(req))



