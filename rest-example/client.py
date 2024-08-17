import requests
import time

start = time.time()

for _ in range(0, 1000):
    response = requests.get('http://localhost:8081/user/1')

end = time.time()

print(f"Last response {response.json()}, elapsed time {end - start}")
