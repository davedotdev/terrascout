from flask import Flask
from flask import request
from tfgenson import SchemaBuilder
import json

app = Flask(__name__)

@app.route('/', methods=['POST'])
def runGenson():
    jsonInput = request.form['payload']
    print(jsonInput)

    builder = SchemaBuilder()
    builder.add_schema({"type": "object", "properties": {}})

    builder.add_object(json.loads(jsonInput), True)

    return builder.to_schema()
