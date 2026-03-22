# Simple Go Server

This is a simple Go server that listens on port 8080

API Endpoints:

- `GET /`: Returns a welcome message "Welcome to the Go Server!".
- `GET /hello`: Returns a JSON message "Hello, World!".
- `GET /form.html`: Returns an HTML form that allows users to submit their name and email.
- `POST /submit`: Accepts form data with `name` and `email` fields and returns a JSON message with the submitted data.