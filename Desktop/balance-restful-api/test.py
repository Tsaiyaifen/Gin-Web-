import requests
import json

base_url = "https://8082-m-s-3dna6vo5z3t52-b.europe-west4-1.prod.colab.dev"

# 存款 10 元
print("=== 存款 10 元 ===")
response = requests.post(base_url + "/deposit/10")
print(f"狀態碼: {response.status_code}")
if response.status_code == 200:
    print("回應:")
    print(json.dumps(response.json(), indent=2, ensure_ascii=False))

# 查詢餘額
print("\n=== 查詢餘額 ===")
response = requests.get(base_url + "/balance/")
print(f"狀態碼: {response.status_code}")
if response.status_code == 200:
    print("回應:")
    print(json.dumps(response.json(), indent=2, ensure_ascii=False))
