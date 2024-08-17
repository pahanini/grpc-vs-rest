# dirty hack
import os
import sys
sys.path.append(os.getcwd() + '/pb')

import time
import grpc
from pb import user_pb2, user_pb2_grpc

channel = grpc.insecure_channel('localhost:50005')
stub = user_pb2_grpc.UserStub(channel)

req = user_pb2.Request(ID=1)
req.ID = 1

start = time.time()

for _ in range(0, 1000):
    response = stub.Info(req)

end = time.time()

print(f"Last response {response}, elapsed time {end - start}")




