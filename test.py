import base64
# import tensorflow as tf
# import PIL
import json
import requests

import base64
with open("CV_API/data/backpack_test.jpg", "rb") as img_file:
    img = base64.b64encode(img_file.read())

requests.post('http://127.0.0.1:5000/get_predictions',
              data=json.dumps(str(img)))
