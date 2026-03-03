from flask import Flask, request, jsonify

app = Flask(__name__)

payments = []
id_counter = 1

@app.route('/payments', methods=['GET'])
def get_payments():
    return jsonify(payments)

@app.route('/payments/process', methods=['POST'])
def process_payment():
    global id_counter
    data = request.json
    payment = {
        "id": id_counter,
        "orderId": data.get("orderId"),
        "amount": data.get("amount"),
        "method": data.get("method"),
        "status": "SUCCESS"
    }
    id_counter += 1
    payments.append(payment)
    return jsonify(payment), 201

@app.route('/payments/<int:id>', methods=['GET'])
def get_payment(id):
    for p in payments:
        if p["id"] == id:
            return jsonify(p)
    return jsonify({"message": "Not found"}), 404

if __name__ == '__main__':
    app.run(host="0.0.0.0", port=8083)