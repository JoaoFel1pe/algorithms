package main

import (
	luhnalgorithm "credit-card-validator/luhn_algorithm"
	"fmt"
)

func main() {
	const express = require('express');
	const app = express();

	app.use(express.json());

	app.get('/validate-card', (req, res) => {
	const {cardNumber} = req.query;
	if (!cardNumber) {
		return res.status(400).json({ error: 'Missing credit card number' });
	}
	const isValid = luhn(cardNumber);
	res.json({isValid});
	});

	app.listen(3000, () => {
	console.log('Server listening on port 3000');
	});

}
