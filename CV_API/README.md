# Running instructions

clone the repository using git to a directory on your local machine and navigate to this directory

Create a virtual environment using `python3 -m venv venv`

Activate the virtual env using `source venv/bin/activate`

Create a shell file named `env_vars.sh` in the `/venv` directory, containing:
```
export FLASK_APP="CV_API/app.py"
export FLASK_ENV="development"
export FLASK_DEBUG= "0"
```

Source this file using `source venv/env_vars.sh`

Install requirements using `pip install -r requirements.txt`

Run the API using `flask run`

Call the endpoint with url `127.0.0.1:5000/get_predictions/{image_num}` where image num is an int between 1&4

Each `image_num` refers to a different image saved locally to get the predictions for