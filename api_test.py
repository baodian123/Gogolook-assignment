import requests

url = "http://localhost:8080/tasks"

headers = {
    "Content-Type": "application/json"
}

def GetTask():
    response = requests.get(url, headers=headers)

    print("Status Code:", response.status_code)
    print("Response Body:", response.text)
    print()

def CreateTask():
    payload = {
        "name": "Test Create",
        "status": 0
    }

    response = requests.post(url, json=payload, headers=headers)

    print("Status Code:", response.status_code)
    print("Response Body:", response.text)
    print()

    try:
        return response.json().get("id")
    except Exception as e:
        print("Failed to parse response JSON:", e)
        return None

def UpdateTask(id):
    payload = {
        "name": "Test Update",
        "status": 1
    }

    response = requests.put(url + "/" + id, json=payload, headers=headers)

    print("Status Code:", response.status_code)
    print("Response Body:", response.text)
    print()

def DeleteTask(id):
    response = requests.delete(url + "/" + id, headers=headers)

    print("Status Code:", response.status_code)
    print("Response Body:", response.text)
    print()

if __name__ == "__main__":
    id = CreateTask()

    GetTask()
    UpdateTask(id)
    GetTask()
    DeleteTask(id)
    GetTask()