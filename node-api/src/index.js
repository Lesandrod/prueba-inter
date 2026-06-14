const express = require("express");
const app = express();
app.use(express.json());

require('./routes')(app);

const PORT = process.env.PORT || 3000;

app.listen(PORT, "0.0.0.0", () => {
  console.log(`Server running on port ${PORT}`);
});