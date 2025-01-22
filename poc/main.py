import asyncio
import websockets
import json
import uuid
import hashlib
import base64
import os

from dotenv import load_dotenv

load_dotenv()


async def get_scene_list():
    uri = "ws://localhost:4455"
    password = os.getenv("OBS_PASSWORD")  # Get the password from the .env file

    try:
        async with websockets.connect(uri) as websocket:

            response = await websocket.recv()
            response_data = json.loads(response)


            if "authentication" in response_data["d"]:

                challenge = response_data["d"]["authentication"]["challenge"]
                salt = response_data["d"]["authentication"]["salt"]

                secret = base64.b64encode(hashlib.sha256((password + salt).encode('utf-8')).digest()).decode('utf-8')
                auth_response = base64.b64encode(hashlib.sha256((secret + challenge).encode('utf-8')).digest()).decode('utf-8')


                await websocket.send(json.dumps({
                    "op": 1,
                    "d": {
                        "rpcVersion": 1,
                        "authentication": auth_response
                    }
                }))

                identified_message = await websocket.recv()
                print("Identified Response:", identified_message)

                identified_data = json.loads(identified_message)
                if identified_data.get("op") != 2:
                    print("Failed to identify with the server.")
                    return
            else:
                await websocket.send(json.dumps({
                    "op": 1,
                    "d": {
                        "rpcVersion": 1,
                    }
                }))
                
                identified_message = await websocket.recv()
                print("Identified Response:", identified_message)
                identified_data = json.loads(identified_message)

                if identified_data.get("op") != 2:
                    print("Failed to identify with the server.")
                    return

            request_id = str(uuid.uuid4())

            await websocket.send(json.dumps({
                "op": 6,
                "d": {
                    "requestType": "GetSceneList",
                    "requestId": request_id
                }
            }))

            response = await websocket.recv()
            print("Scene List Response:", response)

    except Exception as e:
        print(f"Error: {e}")

# Run the async function
asyncio.run(get_scene_list())



