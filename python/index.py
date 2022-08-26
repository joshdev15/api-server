from http.server import BaseHTTPRequestHandler, HTTPServer
from io import StringIO
import json

# Pseudo Router Function
def api(path):
    value = {}
    code = 0

    # TODO: Change the way you manage your routes
    if path == '/':
        value = {'title': 'Simple API Rest Python example', 'success': 'Hola server en Python 3!', 'statusCode': 200}
        code = 200
    elif path == '/error':
        value = {'title': 'Simple API Rest Python example', 'error': 'Esto es un error!', 'statusCode': 500}
        code = 500

    return value, code
        
# Server Class
class Server(BaseHTTPRequestHandler):
    def do_GET(self):
        # Pseudorouter
        response, status = api(self.path)

        # Abstract Headers
        self.send_response(status)
        self.send_header("Content-type", "application/json")
        self.end_headers()

        # Formatting JSON response
        customJSON = json.dumps(response)
        self.wfile.write(bytes(customJSON,'utf-8'))


# Run Server Function
def runServer():
    hostname = "localhost"
    port = 8080
    webServer = HTTPServer((hostname, port), Server)
    print("Your server is running! - http://%s:%s" % (hostname, port))
    webServer.serve_forever()

# Run Server Call
runServer()

# Made in Python 3.10.6
