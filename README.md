# Fetch Backend Engineering Challenge
## Receipt Processor
Build a webservice that fulfils the documented API. 

### Endpoint: Process Receipts
Path: /receipts/process

Method: POST

Payload: Receipt JSON

Response: JSON containing an id for the receipt.

### Endpoint: Get Points
Path: /receipts/{id}/points

Method: GET

Response: A JSON object containing the number of points awarded.

### How to Run

Use the following commands:
1. docker build -t receipt-processor .
2. docker run -p 8080:8080 receipt-processor