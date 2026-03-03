const express = require('express');
const app = express();
app.use(express.json());

let items = [
  { id: 1, name: "Book" },
  { id: 2, name: "Laptop" }
];

let idCounter = 3;

app.get('/items', (req, res) => {
  res.json(items);
});

app.post('/items', (req, res) => {
  const newItem = {
    id: idCounter++,
    name: req.body.name
  };
  items.push(newItem);
  res.status(201).json(newItem);
});

app.get('/items/:id', (req, res) => {
  const item = items.find(i => i.id == req.params.id);
  if (!item) return res.status(404).send("Not found");
  res.json(item);
});

app.listen(8081, () => {
  console.log("Item Service running on port 8081");
});