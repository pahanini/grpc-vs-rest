import requests

response = requests.get('http://localhost:8081/user/1')

print(response.json())
